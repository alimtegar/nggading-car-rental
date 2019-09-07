package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty" validate:"nonzero"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"nonzero"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"nonzero"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"nonzero"`
}
