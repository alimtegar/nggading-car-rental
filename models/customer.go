package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer Model
type Customer struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"username,omitempty" validate:"nonzero"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty" validate:"nonzero"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"nonzero"`
	NIK        string             `json:"nik,omitempty" bson:"nik,omitempty" validate:"nonzero"`
	STNKNumber string             `json:"stnkNumber,omitempty" bson:"stnkNumber,omitempty" validate:"nonzero"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty" validate:"nonzero"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}
