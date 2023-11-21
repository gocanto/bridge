package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Webhook Redirector",
		IdleTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
		Prefork:     true,
	})

	app.Group("webhooks", func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		app.Post("/handle", func(c *fiber.Ctx) error {
			return c.SendString("here!")
		})

		return nil
	})

	if result := app.Listen("localhost:8081"); result != nil {
		fmt.Println("error ...")
	}
}
