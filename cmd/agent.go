package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/samkreter/kube-network-watcher/agent/apiserver"
	agentService "github.com/samkreter/kube-network-watcher/proto/agent/v1"
)


var (
	agentGRPCAddr string
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "The agent to perform the network watchs",
	Long:  `The agent to perform the network watch`,
	Run:   agentRun,
}

func init() {
	rootCmd.AddCommand(agentCmd)
	agentCmd.Flags().StringVar(&agentGRPCAddr, "controllerGRPCAddr", ":5656", "thee addr for the controller to server its grpc server")
}

func agentRun(cmd *cobra.Command, args []string) {
	svr := apiserver.NewServer()

	listen, err := net.Listen("tcp", agentGRPCAddr)
	if err != nil {
		log.Fatalf("Error attaching TPC listener to address %s: %s", agentGRPCAddr, err.Error())
	}

	grpcServer := grpc.NewServer()
	agentService.RegisterAgentServiceServer(grpcServer, svr)
	reflection.Register(grpcServer)

	errorPipeline := make(chan error)
	go func(){
		log.Println("starting grpc agent server")
		errorPipeline <- grpcServer.Serve(listen)
	}()

	if err := <-errorPipeline; err != nil {
		log.Fatal(err)
	}
}

