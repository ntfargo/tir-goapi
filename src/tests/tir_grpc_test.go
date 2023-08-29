package tests

import (
	"context"
	"log"
	"time"

	tirengine "github.com/ntfargo/tir-goapi/src/tirengine" // generated from tir.proto
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := tirengine.NewTirServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.GenerateKnowledge(ctx, &tirengine.EmptyRequest{})
	if err != nil {
		log.Fatalf("Failed to call GenerateKnowledge at %s: %v", err)
	} else {
		log.Printf("Called GenerateKnowledge successfully at %s!")
	}
}
