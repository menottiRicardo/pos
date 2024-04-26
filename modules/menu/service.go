package menu

import (
	"context"
	"fmt"
	"pos/pkg/common"
	"pos/pkg/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateItem creates a new menu item in the database.
func CreateItem(c *fiber.Ctx) (*MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var menuItem MenuItem
	if err := c.BodyParser(&menuItem); err != nil {
		return nil, err
	}
	collection := db.OpenCollection(db.Client, "menu", "items")

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

	collection := db.OpenCollection(db.Client, "menu", "items")

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

	collection := db.OpenCollection(db.Client, "menu", "items")
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

	collection := db.OpenCollection(db.Client, "menu", "items")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": itemObjID})
	if err != nil {
		return "", err
	}

	return itemID, nil
}

// // // Categories // // //
func CreateCategory(c *fiber.Ctx) (*MenuCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var category MenuCategory
	if err := c.BodyParser(&category); err != nil {
		return nil, err
	}

	validator := common.NewValidator()
	// Validate the menuItem struct
	if err := validator.Validate(category); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	category.ID = primitive.NewObjectID()

	collection := db.OpenCollection(db.Client, "menu", "categories")
	_, err := collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategory(categoryID string) (*MenuCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var category MenuCategory
	catObjID, _ := primitive.ObjectIDFromHex(categoryID)
	collection := db.OpenCollection(db.Client, "menu", "categories")
	err := collection.FindOne(ctx, bson.M{"_id": catObjID}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory updates information about an existing category.
func UpdateCategory(categoryID string, updatedCategory MenuCategory) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	catObjID, _ := primitive.ObjectIDFromHex(categoryID)

	update := bson.M{"$set": updatedCategory}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	collection := db.OpenCollection(db.Client, "menu", "categories")
	var result mongo.UpdateResult
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": catObjID}, update, opts).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCategory removes a category from the database.
func DeleteCategory(categoryID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	catObjID, _ := primitive.ObjectIDFromHex(categoryID)
	collection := db.OpenCollection(db.Client, "menu", "categories")
	// TODO: Add functionality to delete all items from that category
	_, err := collection.DeleteOne(ctx, bson.M{"_id": catObjID})
	if err != nil {
		return "", err
	}
	return categoryID, nil
}

// menu

func CreateMenu(c *fiber.Ctx) (*Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var menu Menu
	if err := c.BodyParser(&menu); err != nil {
		return nil, err
	}

	validator := common.NewValidator()
	// Validate the menuItem struct
	if err := validator.Validate(menu); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	menu.ID = primitive.NewObjectID()

	collection := db.OpenCollection(db.Client, "menu", "menus")
	_, err := collection.InsertOne(ctx, menu)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func UpdateMenu(menuID string, updatedMenu Menu) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	catObjID, _ := primitive.ObjectIDFromHex(menuID)

	update := bson.M{"$set": updatedMenu}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	collection := db.OpenCollection(db.Client, "menu", "categories")
	var result mongo.UpdateResult
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": catObjID}, update, opts).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetCompleteMenu(menuID string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	menuObjID, _ := primitive.ObjectIDFromHex(menuID)
	collection := db.OpenCollection(db.Client, "menu", "menus")

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", menuObjID}}}},
		{{"$lookup", bson.D{
			{"from", "categories"},
			{"localField", "_id"},
			{"foreignField", "menuId"},
			{"as", "categories"},
		}}},
		{{"$unwind", "$categories"}},
		{{"$project", bson.D{
			{"categories.itemsIds", 0}, // Omit the description field from items
		}}},
		{{"$lookup", bson.D{
			{"from", "items"},
			{"localField", "categories.itemIds"},
			{"foreignField", "_id"},
			{"as", "categories.items"},
		}}},
		{{"$project", bson.D{
			{"categories.itemIds", 0}, // Omit the itemIds field
		}}},
		{{"$group", bson.D{
			{"_id", "$_id"},
			{"name", bson.M{"$first": "$name"}},
			{"categories", bson.M{"$push": "$categories"}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregation error: %w", err)
	}

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("reading cursor error: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no menu found with ID %s", menuID)
	}

	return results, nil
}
