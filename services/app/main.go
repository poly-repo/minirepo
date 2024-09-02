package main

import (
	"log"

	"go/test/services/lib/helloworld"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(helloworld.Hello("whole World"))
	})

	log.Fatal(app.Listen(":3000"))
}
