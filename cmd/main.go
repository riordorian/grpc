package main

import (
	"context"
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	"grpc/dependencies"
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
	c := len(dependencies.ConfigServices) +
		len(dependencies.ConfigServices) +
		len(dependencies.RepositoryServices) +
		len(dependencies.AuthServices) +
		len(dependencies.ApplicationServices) +
		len(dependencies.PortsServices)

	var defs = make([]di.Def, 0, c)
	defs = append(defs, dependencies.ConfigServices...)
	defs = append(defs, dependencies.RepositoryServices...)
	defs = append(defs, dependencies.AuthServices...)
	defs = append(defs, dependencies.ApplicationServices...)
	defs = append(defs, dependencies.PortsServices...)
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
