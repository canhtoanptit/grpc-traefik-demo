package main

import (
	"context"
	"log"
	"time"

	"github.com/canhtoanptit/traefik-grpc-example/card-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewCardServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateCard(ctx, &proto.CreateCardRequest{
		CardNumber: "1234-5678-9012-3456",
		OwnerName:  "John Doe",
	})
	if err != nil {
		log.Fatalf("Failed to call CreateCard: %v", err)
	}

	log.Printf("Card ID: %s, Message: %s", resp.CardId, resp.Message)
}
