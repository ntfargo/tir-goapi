package tests

import (
	"context"
	"testing"
	"time"

	tir "github.com/ntfargo/tir-goapi/src/tir-engine/proto" // generated from tir.proto
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestGenerateKnowledge(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := tir.NewTirServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.GenerateKnowledge(ctx, &tir.GenerateKnowledgeRequest{})
	if err != nil {
		t.Fatalf("Failed to call GenerateKnowledge: %v", err)
	}
}
