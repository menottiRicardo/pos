package menu

import "go.mongodb.org/mongo-driver/bson/primitive"

// MenuItem represents an item in a menu category
type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description" validate:"required"`
	Price       float64            `bson:"price" json:"price" validate:"gte=0"`
	Available   bool               `bson:"available" json:"available"`
	Ingredients []string           `bson:"ingredients" json:"ingredients"`
}

// MenuCategory represents a category in the menu
type MenuCategory struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name" validate:"required"`
	Items []MenuItem         `bson:"items" json:"items"`
}

// Menu represents the complete menu structure
type Menu struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	Categories []MenuCategory     `bson:"categories" json:"categories"`
}
