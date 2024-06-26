package main

import (
	"context"
	"errors"
	"log"

	"pos/modules/menu"
	"pos/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			return c.Status(code).SendString("something happened at the disco!")
		},
	})

	// Middleware to log the HTTP requests
	app.Use(logger.New())

	app.Use(recover.New())

	// Setup MongoDB connection
	client := db.Client

	// connect to mongodb
	//Client := database.ConnectWithMongodb()

	// close the connection
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.Background())

	// Setup routes
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {

	// menu items

	app.Put("/menu/items/:id", menu.UpdateMenuItem)
	app.Post("/menu/items", menu.CreateMenuItem)
	app.Get("/menu/items/:id", menu.GetMenuItem)
	app.Delete("/menu/items/:id", menu.DeleteMenuItem)

	// categories
	app.Post("/menu/categories", menu.CreateCategoryHandler)
	app.Put("/menu/categories/:id", menu.UpdateCategoryHandler)
	app.Get("/menu/categories/:id", menu.GetCategoryHandler)

	// menu
	app.Post("/menu", menu.CreateMenuHandler)
	app.Put("/menu/:id", menu.UpdateMenuHandler)
	app.Get("/menu/:id", menu.GetMenuHandler)
}
