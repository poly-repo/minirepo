package main

import (
	"fmt"
	"log"

	portal_pb "go/test/proto"
	"uno/services/lib/helloworld"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
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
