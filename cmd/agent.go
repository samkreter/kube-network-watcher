package cmd

import (
	"net"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	log "github.com/sirupsen/logrus"

	"github.com/samkreter/kube-network-watcher/agent/apiserver"
	"github.com/samkreter/kube-network-watcher/tcpdump"
	agentService "github.com/samkreter/kube-network-watcher/proto/agent/v1"
)


var (
	agentGRPCAddr string
	defaultPodName = "local"
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
	podName := os.Getenv("POD_NAME")
	deployName := os.Getenv("DEPLOYMENT_NAME")
	namespace := os.Getenv("NAMESPACE")

	svr := apiserver.NewServer(podName, deployName , namespace )

	listen, err := net.Listen("tcp", agentGRPCAddr)
	if err != nil {
		log.Fatalf("Error attaching TPC listener to address %s: %s", agentGRPCAddr, err.Error())
	}

	grpcServer := grpc.NewServer()
	agentService.RegisterAgentServiceServer(grpcServer, svr)
	reflection.Register(grpcServer)

	errorPipeline := make(chan error)
	go func(){
		log.Info("starting grpc agent server")
		errorPipeline <- grpcServer.Serve(listen)
	}()

	go func(){
		log.Info("starting tcp dumping")
		if err := tcpdump.StartPacketDump(); err != nil {
			errorPipeline <- err
		}
	}()

	if err := <-errorPipeline; err != nil {
		log.Fatal(err)
	}
}


