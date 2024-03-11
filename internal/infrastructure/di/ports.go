package di

import (
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	gpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports"
	"grpc/internal/infrastructure/ports/grpc"
	"grpc/internal/infrastructure/ports/grpc/convertors"
	"grpc/internal/infrastructure/ports/grpc/interceptors"
	"grpc/internal/shared/interfaces"
	gproto "grpc/pkg/proto_gen/grpc"
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
			authInterceptor := ctn.Get("AuthInterceptor").(interceptors.AuthInterceptor)
			server := gpc.NewServer(gpc.ChainUnaryInterceptor(authInterceptor.Get()))
			s := &grpc.NewsServer{
				Server:   server,
				Listener: lis,
				Handlers: ctn.Get("ApplicationHandlers").(application.Handlers),
				Convertors: grpc.Convertors{
					ListRequest:   new(convertors.ListRequestConvertor),
					ListResponse:  new(convertors.ListResponseConvertor),
					LoginRequest:  new(convertors.UserLoginRequestConvertor),
					LoginResponse: new(convertors.UserLoginResponseConvertor),
				},
			}

			reflection.Register(s.Server)
			//pg.RegisterNewsServer(s.Server, NewsServer{})
			gproto.RegisterNewsServer(s.Server, s)
			gproto.RegisterUsersServer(s.Server, s)
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
	{
		Name:  "AuthInterceptor",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return interceptors.AuthInterceptor{
				AuthProvider: ctn.Get("AuthProvider").(interfaces.AuthProviderInterface),
			}, nil
		},
	},
}
