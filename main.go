package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	fmt.Println("gus")

	app := fiber.New(fiber.Config{
		AppName:     "Webhook Redirector",
		IdleTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
		Prefork:     true,
	})

	app.Post("/handle", func(c *fiber.Ctx) error {
		return c.SendString("here!")
	})

	fmt.Println("listening ...")

	if result := app.Listen("localhost:8081"); result != nil {
		fmt.Println("error ...")
	}
}
