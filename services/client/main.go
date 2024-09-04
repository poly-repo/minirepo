package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	portal_pb "uno/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	target := "demo.grpc.polymath-solutions.com:443"
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, // This is equivalent to the -insecure flag in grpcurl
		})),
	}
	//	target := "localhost:50053"
	//	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(target, opts...)

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
