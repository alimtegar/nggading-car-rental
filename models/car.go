package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand       string             `json:"brand,omitempty" bson:"brand,omitempty" validate:"nonzero"`
	Model       string             `json:"model,omitempty" bson:"model,omitempty" validate:"nonzero"`
	Year        int32              `json:"year,omitempty" bson:"year,omitempty" validate:"nonzero"`
	Color       string             `json:"color,omitempty" bson:"color,omitempty" validate:"nonzero"`
	PlatNumber  string             `json:"platNumber,omitempty" bson:"platNumber,omitempty" validate:"nonzero"`
	Stock       int32              `json:"stock,omitempty" bson:"stock,omitempty" validate:"nonzero"`
	Price       int32              `json:"price,omitempty" bson:"price,omitempty" validate:"nonzero"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" validate:"nonzero"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}