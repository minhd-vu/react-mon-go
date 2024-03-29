package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model schema
type User struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
