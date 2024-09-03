package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"time"

	portal_pb "uno/proto"
	"uno/services/lib/helloworld"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
}

func (s *server) Eval(ctx context.Context, in *portal_pb.EvalRequest) (*portal_pb.EvalResponse, error) {
	return &portal_pb.EvalResponse{Result: "Hello " + in.Expr}, nil
}

func main() {
	go func() {
		w := os.Stderr
		logger := slog.New(tint.NewHandler(w, nil))

		// set global logger with custom options
		slog.SetDefault(slog.New(
			tint.NewHandler(w, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		))
		logger.Info("Something weird")

		app := fiber.New()

		msg := portal_pb.EvalRequest{Expr: "Bar"}
		fmt.Println(protojson.Format(&msg))
		newUUID := uuid.New()
		fmt.Printf("UUID: %v\n", newUUID)

		app.Get("/", func(c *fiber.Ctx) {
			c.SendString(helloworld.Hello("whole World"))
		})

		log.Fatal(app.Listen(":3000"))
	}()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	portal_pb.RegisterEvaluatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
