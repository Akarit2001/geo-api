package entities

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Description string             `json:"description"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	Metadata    *json.RawMessage   `json:"meta,omitempty" bson:"meta,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

type Properties struct {
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	CollectionID primitive.ObjectID `json:"collectionId,omitempty" bson:"collectionId,omitempty"`
}

type Feature struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type,omitempty" bson:"type,omitempty"`
	Geometry   Geometry           `json:"geometry,omitempty" bson:"geometry,omitempty"`
	Properties Properties         `json:"properties,omitempty" bson:"properties,omitempty"`
}
