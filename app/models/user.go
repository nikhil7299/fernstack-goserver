package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone           string             `json:"phone,omitempty" bson:"phone,omitempty"`
	IsPhoneVerified bool               `json:"isPhoneVerified,omitempty" bson:"isPhoneVerified,omitempty"`
	Type            string             `json:"type,omitempty" bson:"type,omitempty" default:"user"`
}
