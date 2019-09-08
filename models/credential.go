package models

// Credential Model
type Credential struct {
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"nonzero"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"nonzero"`
}
