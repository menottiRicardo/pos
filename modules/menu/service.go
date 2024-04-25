package menu

import (
	"context"
	"pos/pkg/common"
	"pos/pkg/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = db.OpenCollection(db.Client, "menuItems")

// CreateItem creates a new menu item in the database.
func CreateItem(c *fiber.Ctx) (*MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var menuItem MenuItem
	if err := c.BodyParser(&menuItem); err != nil {
		return nil, err
	}

	validator := common.NewValidator()
	// Validate the menuItem struct
	if err := validator.Validate(menuItem); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	menuItem.ID = primitive.NewObjectID()

	// Insert the new menu item into the database
	_, err := collection.InsertOne(ctx, menuItem)
	if err != nil {
		return nil, err
	}

	return &menuItem, nil
}
