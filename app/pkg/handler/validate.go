package handler

import (
	"context"
	"log"

	pb "github.com/fxproxy/app/pkg/proto"
)

func (s *Server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("app service is working for SayHello...Received rpc from sidecar")
	return &pb.SayHelloResponse{
		Message: "hello " + req.GetMessage(),
	}, nil
}
