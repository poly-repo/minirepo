package main

import (
	"context"
	"fmt"
	"log"

	portal_pb "uno/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	target := "demo.polymath-solutions.com:50053"
	//	target := "localhost:50053"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer conn.Close()

	client := portal_pb.NewEvaluatorClient(conn)

	request := &portal_pb.EvalRequest{
		Expr: "World!",
	}
	resp, err := client.Eval(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Response: %v\n", resp)
}
