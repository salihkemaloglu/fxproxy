package handler

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/fxproxy/sidecar/pkg/proto"
	"github.com/fxproxy/sidecar/pkg/util/rpc"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SayHello(ctx *gin.Context) {
	resp := []string{"sidecar is working ..."}

	client, conn, err := rpc.App()
	if err != nil {
		resp = append(resp, fmt.Sprintf("could not dial account service: %v", err))
		ctx.JSON(http.StatusGatewayTimeout, resp)
		return
	}
	defer conn.Close()

	r, err := client.SayHello(context.Background(), &pb.SayHelloRequest{Message: "hello"})
	if err != nil {
		resp = append(resp, fmt.Sprintf("could not response from SayHello: %v", err))
		ctx.JSON(http.StatusGatewayTimeout, resp)
		return
	}

	resp = append(resp, r.GetMessage())
	ctx.JSON(http.StatusOK, resp)
}
