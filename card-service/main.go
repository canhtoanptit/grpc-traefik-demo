package main

import (
	"context"
	"log"
	"net"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"github.com/canhtoanptit/traefik-grpc-example/card-service/proto"
)

type server struct {
	proto.UnimplementedCardServiceServer
}

func (s *server) CreateCard(ctx context.Context, req *proto.CreateCardRequest) (*proto.CreateCardResponse, error) {
	cardID := uuid.New().String()
	return &proto.CreateCardResponse{
		CardId:  cardID,
		Message: "Card created for " + req.OwnerName + " with number " + req.CardNumber,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCardServiceServer(s, &server{})

	log.Println("Card gRPC server running on :50052")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}