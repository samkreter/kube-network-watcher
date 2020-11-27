package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"


	controllerService "github.com/samkreter/kube-network-watcher/proto/controller/v1"
	"github.com/samkreter/kube-network-watcher/controller/apiserver"
)


var (
	controllerGRPCAddr string
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
	Run:   controllerRun,
}

func init() {
	rootCmd.AddCommand(controllerCmd)
	controllerCmd.Flags().StringVar(&controllerGRPCAddr, "controllerGRPCAddr", ":5656", "thee addr for the controller to server its grpc server")
}

func controllerRun(cmd *cobra.Command, args []string) {
	svr := apiserver.NewServer()

	listen, err := net.Listen("tcp", controllerGRPCAddr)
	if err != nil {
		log.Fatalf("Error attaching TPC listener to address %s: %s", controllerGRPCAddr, err.Error())
	}

	grpcServer := grpc.NewServer()
	controllerService.RegisterControllerServiceServer(grpcServer, svr)
	reflection.Register(grpcServer)

	errorPipeline := make(chan error)
	go func(){
		log.Println("starting grpc controller server")
		errorPipeline <- grpcServer.Serve(listen)
	}()

	if err := <-errorPipeline; err != nil {
		log.Fatal(err)
	}
}

