package dependencies

import (
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports"
	"grpc/internal/infrastructure/ports/grpc"
	"grpc/internal/infrastructure/ports/grpc/convertors"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"log"
	"net"
)

var PortsServices = []di.Def{
	{
		Name:  "GrpcServer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("ConfigProvider").(*viper.Viper)
			address := fmt.Sprintf("%s:%s", config.GetString("GRPC_SERVER_HOST"), config.GetString("GRPC_SERVER_PORT"))
			lis, err := net.Listen("tcp", address)
			// TODO: Add logger
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			server := gp.NewServer(gp.UnaryInterceptor(grpc.UnaryServerInterceptor))
			s := &grpc.NewsServer{
				Server:   server,
				Listener: lis,
				Handlers: ctn.Get("ApplicationHandlers").(application.Handlers),
				Convertors: grpc.Convertors{
					ListRequest:  new(convertors.ListRequestConvertor),
					ListResponse: new(convertors.ListResponseConvertor),
					LoginRequest: new(convertors.UserLoginRequestConvertor),
				},
			}

			reflection.Register(s.Server)
			//pg.RegisterNewsServer(s.Server, NewsServer{})
			pg.RegisterNewsServer(s.Server, s)
			pg.RegisterUsersServer(s.Server, s)
			//pg.RegisterPromosServer(server, pg.UnimplementedPromosServer{})

			return s, nil

		},
	},
	{
		Name:  "PortsServices",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return ports.Services{
				GrpcServer: ctn.Get("GrpcServer").(*grpc.NewsServer),
				//HttpServer: http.GetServer(appServices.Handler),
			}, nil
		},
	},
}
