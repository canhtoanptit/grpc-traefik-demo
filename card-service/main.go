package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/canhtoanptit/traefik-grpc-example/card-service/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
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
	// Start gRPC server on :8081
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen gRPC: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCardServiceServer(s, &server{})
	log.Println("Card gRPC server running on :8081")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server on :9090
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Card Service Web Interface"))
	})
	log.Println("Card HTTP server running on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
