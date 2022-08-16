package greeting

import (
	"context"
)

type grpcServer struct{}

func NewGRPCServer() *grpcServer {
	return &grpcServer{}
}

func (s *grpcServer) SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello again"}, nil
}

func (s *grpcServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello"}, nil
}

func (s *grpcServer) mustEmbedUnimplementedGreeterServer() {}
