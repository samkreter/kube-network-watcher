package apiserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	types "github.com/samkreter/kube-network-watcher/proto/agent/v1"
)


type Server struct {
	PodName string
	DeployName string
	Namespace string
}

func NewServer(podName, deployName, namespace string) *Server {
	return &Server{
		PodName: podName,
		DeployName: deployName,
		Namespace: namespace,
	}
}

func (s *Server) StartNetworkWatch(ctx context.Context, startNetworkWatchRequest *types.StartNetworkWatchRequest) (*types.StartNetworkWatchResponse, error){
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
func (s *Server) StopNetworkWatch(ctx context.Context, stopNetworkWatchRequest *types.StopNetworkWatchRequest) (*types.StopNetworkWatchResponse, error){
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

