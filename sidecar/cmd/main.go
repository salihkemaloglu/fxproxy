package main

import (
	"github.com/fxproxy/sidecar/pkg/handler"
	"github.com/fxproxy/sidecar/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handler.NewHandler()
	r.Use(middleware.ValidatePathMiddleware())
	r.GET("/hello", h.SayHello)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
