package main

import (
	"fmt"
	"grpc-gopher-lets-go/message"
	"grpc-gopher-lets-go/service"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("Failed to listen: %+v", err)
	}
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	srv := service.NewMessageServer()
	message.RegisterMessageServiceServer(s, srv)
	fmt.Println("Start listening localhost:10000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
