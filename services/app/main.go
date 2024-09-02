package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	portal_pb "go/test/proto"
	"uno/services/lib/helloworld"

	"github.com/google/uuid"
	"github.com/lmittmann/tint"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	w := os.Stderr
	logger := slog.New(tint.NewHandler(w, nil))

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
	logger.Error("Something weird")

	app := fiber.New()

	msg := portal_pb.EvalRequest{Name: "Foo"}
	fmt.Println(protojson.Format(&msg))
	newUUID := uuid.New()
	fmt.Printf("UUID: %v\n", newUUID)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(helloworld.Hello("whole World"))
	})

	log.Fatal(app.Listen(":3000"))
}
