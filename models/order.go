package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order Model
type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `json:"customer_id,omitempty" bson:"customer_id,omitempty" validate:"nonzero"`
	CarID      primitive.ObjectID `json:"car_id,omitempty" bson:"car_id,omitempty" validate:"nonzero"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}

// RequestedOrder Model
type RequestedOrder struct {
	Name       string `json:"name,omitempty" bson:"username,omitempty" validate:"nonzero"`
	Email      string `json:"email,omitempty" bson:"email,omitempty" validate:"nonzero"`
	Phone      string `json:"phone,omitempty" bson:"phone,omitempty" validate:"nonzero"`
	NIK        string `json:"nik,omitempty" bson:"nik,omitempty" validate:"nonzero"`
	STNKNumber string `json:"stnkNumber,omitempty" bson:"stnkNumber,omitempty" validate:"nonzero"`
	Address    string `json:"address,omitempty" bson:"address,omitempty" validate:"nonzero"`
	CarID      string `json:"car_id,omitempty" bson:"car_id,omitempty" validate:"nonzero"`
}
