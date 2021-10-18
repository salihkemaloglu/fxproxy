package rpc

import (
	pb "github.com/fxproxy/sidecar/pkg/proto"
	"google.golang.org/grpc"
)

//App Dial to account service
func App() (pb.AppServiceClient, *grpc.ClientConn, error) {
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50001", opts)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewAppServiceClient(conn), conn, nil
}
