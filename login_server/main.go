//go:generate protoc -I ../grpc_learn --go_out=plugins=grpc:$GOPATH/src ../model/model.proto
//go:generate protoc -I ../grpc_learn --go_out=plugins=grpc:$GOPATH/src ../service/login.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/imfuxiao/grpc_learn/model"
	"github.com/imfuxiao/grpc_learn/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type loginServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *loginServer) Login(ctx context.Context, in *model.LoginRequest) (*model.UserInfo, error) {
	log.Printf("Received: %+v", in)
	return &model.UserInfo{
		Name: in.Account,
	}, nil;
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterLoginServiceServer(s, &loginServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
