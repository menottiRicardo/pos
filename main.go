package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // Ensure this import is correct
)

func main() {
	app := fiber.New()

	// Middleware to log the HTTP requests
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	log.Fatal(app.Listen(":3000"))
}
