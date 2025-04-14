package main

import (
	"context"
	"log"
	"net"

	"github.com/canhtoanptit/traefik-grpc-example/go-app/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	log.Println("gRPC server running on :8082")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
