package main

import (
	"fmt"
	"net"

	"github.com/fxproxy/app/pkg/handler"

	pb "github.com/fxproxy/app/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	port := ":50001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{}

	server := handler.NewServer()
	s := grpc.NewServer(opts...)
	pb.RegisterAppServiceServer(s, server)

	fmt.Printf("Account service started as http and listen to port: %v \n", port)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
