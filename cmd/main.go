package main

import (
	"context"
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	dic "grpc/internal/infrastructure/di"
	"grpc/internal/infrastructure/ports"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Grpc server commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start grpc server",
	Long:  `Start grpc server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside serve PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		_, cancel := context.WithCancel(context.Background())
		defer cancel()

		app := bootstrap()
		defer destroy(app)
		portsServices := app.Get("PortsServices").(ports.Services)

		portsServices.GrpcServer.Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside serve PostRun with args: %v\n", args)
	},
}

func bootstrap() di.Container {

	builder, builderErr := di.NewBuilder()
	if builderErr != nil {
		panic(builderErr.Error())
	}
	c := len(dic.ConfigServices) +
		len(dic.ConfigServices) +
		len(dic.RepositoryServices) +
		len(dic.AuthServices) +
		len(dic.SearchServices) +
		len(dic.ApplicationServices) +
		len(dic.PortsServices) +
		len(dic.FileStorage)

	var defs = make([]di.Def, 0, c)
	defs = append(defs, dic.ConfigServices...)
	defs = append(defs, dic.RepositoryServices...)
	defs = append(defs, dic.AuthServices...)
	defs = append(defs, dic.SearchServices...)
	defs = append(defs, dic.ApplicationServices...)
	defs = append(defs, dic.PortsServices...)
	defs = append(defs, dic.FileStorage...)

	builderErr = builder.Add(defs...)
	if builderErr != nil {
		log.Fatal(builderErr.Error())
	}

	app := builder.Build()

	return app
}

func destroy(app di.Container) {
	diDeleteErr := app.Delete()
	if diDeleteErr != nil {
		panic(diDeleteErr.Error())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.Execute()
}
