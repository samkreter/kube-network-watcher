// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	StartNetworkWatch(ctx context.Context, in *StartNetworkWatchRequest, opts ...grpc.CallOption) (*StartNetworkWatchResponse, error)
	StopNetworkWatch(ctx context.Context, in *StopNetworkWatchRequest, opts ...grpc.CallOption) (*StopNetworkWatchResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) StartNetworkWatch(ctx context.Context, in *StartNetworkWatchRequest, opts ...grpc.CallOption) (*StartNetworkWatchResponse, error) {
	out := new(StartNetworkWatchResponse)
	err := c.cc.Invoke(ctx, "/agent.v1.AgentService/StartNetworkWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StopNetworkWatch(ctx context.Context, in *StopNetworkWatchRequest, opts ...grpc.CallOption) (*StopNetworkWatchResponse, error) {
	out := new(StopNetworkWatchResponse)
	err := c.cc.Invoke(ctx, "/agent.v1.AgentService/StopNetworkWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	StartNetworkWatch(context.Context, *StartNetworkWatchRequest) (*StartNetworkWatchResponse, error)
	StopNetworkWatch(context.Context, *StopNetworkWatchRequest) (*StopNetworkWatchResponse, error)
}

// UnimplementedAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) StartNetworkWatch(context.Context, *StartNetworkWatchRequest) (*StartNetworkWatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNetworkWatch not implemented")
}
func (*UnimplementedAgentServiceServer) StopNetworkWatch(context.Context, *StopNetworkWatchRequest) (*StopNetworkWatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopNetworkWatch not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_StartNetworkWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNetworkWatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StartNetworkWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.v1.AgentService/StartNetworkWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StartNetworkWatch(ctx, req.(*StartNetworkWatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StopNetworkWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNetworkWatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StopNetworkWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.v1.AgentService/StopNetworkWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StopNetworkWatch(ctx, req.(*StopNetworkWatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.v1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartNetworkWatch",
			Handler:    _AgentService_StartNetworkWatch_Handler,
		},
		{
			MethodName: "StopNetworkWatch",
			Handler:    _AgentService_StopNetworkWatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/v1/agent.proto",
}
