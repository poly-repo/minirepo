package main

import (
	"fmt"
	"log"

	"go/test/services/lib/helloworld"

	portal_pb "go/test/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := fiber.New()

	msg := portal_pb.EvalRequest{Name: "Foo"}
	fmt.Println(protojson.Format(&msg))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(helloworld.Hello("whole World"))
	})

	log.Fatal(app.Listen(":3000"))
}
