package menu

import (
	"context"
	"pos/pkg/common"
	"pos/pkg/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = db.OpenCollection(db.Client, "menu", "menuItems")

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

// GetItem retrieves a single menu item by its ID.
func GetItem(itemID string) (*MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var item MenuItem
	itemObjID, _ := primitive.ObjectIDFromHex(itemID)
	err := collection.FindOne(ctx, bson.M{"_id": itemObjID}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// UpdateItem updates an existing menu item.
func UpdateItem(itemID string, updatedItem MenuItem) (*MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	itemObjID, _ := primitive.ObjectIDFromHex(itemID)

	update := bson.M{"$set": updatedItem}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var result MenuItem
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": itemObjID}, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteItem deletes a menu item from the database.
func DeleteItem(itemID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	itemObjID, _ := primitive.ObjectIDFromHex(itemID)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": itemObjID})
	if err != nil {
		return "", err
	}

	return itemID, nil
}
