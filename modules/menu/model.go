package menu

import "go.mongodb.org/mongo-driver/bson/primitive"

type MenuItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required,min=2,max=100"`
	Description string             `json:"description" validate:"max=255"`
	Price       float64            `json:"price" validate:"required,gt=0"`
	Available   bool               `json:"available"`
	Category    string             `json:"category" validate:"required,oneof=starter main dessert"`
}
