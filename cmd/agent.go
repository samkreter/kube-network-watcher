package main

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"


	agentService "github.com/samkreter/kube-network-watcher/proto/agent/v1"
	"github.com/samkreter/kube-network-watcher/agent/apiserver"
)


var (
	agentGRPCAddr string
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
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

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

