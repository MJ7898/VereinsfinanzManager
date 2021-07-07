package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NonHumanResources struct {
	SchemaVersion string             `json:"schema_version" bson:"schema_version,omitempty"`
	Key           string             `json:"key,omitempty"`
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Cost          float64            `json:"cost" bson:"cost,omitempty"`
	Validity      string             `json:"validity" bson:"validity,omitempty"`
	TimeStamp     string             `json:"duration" bson:"duration,omitempty"`
	TeamID        primitive.ObjectID `json:"team_id" bson:"team_id"`
}
