package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Player struct
type Player struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
