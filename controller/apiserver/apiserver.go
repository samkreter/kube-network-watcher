package apiserver

import (
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	types "github.com/samkreter/kube-network-watcher/proto/controller/v1"
)



type Server struct {}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Register(ctx context.Context, registerRequest *types.RegisterRequest) (*types.RegisterResponse, error){
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
func (s *Server) DeRegister(ctx context.Context, deRegisterRequest *types.DeRegisterRequest) (*types.DeRegisterResponse, error){
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
