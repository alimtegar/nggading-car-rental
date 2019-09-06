package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `json:"customer,omitempty" bson:"customer,omitempty" validate:"nonzero"`
	Car        primitive.ObjectID `json:"car,omitempty" bson:"car,omitempty" validate:"nonzero"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}