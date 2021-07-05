package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NonHumanResources struct {
	SchemaVersion string `json:"schema_version" bson:"schema_version,omitempty"`
	Key string `json:"key,omitempty"`
	ID     primitive.ObjectID  `bson:"_id,omitempty"`
	Name         string       `json:"name" bson:"name,omitempty"`
	Cost         string       `json:"cost" bson:"cost,omitempty"`
	Validity     string       `json:"validity" bson:"validity,omitempty"`
	TimeStamp    time.Duration    `json:"duration" bson:"duration,omitempty"`
	TeamID primitive.ObjectID `json:"team_id"`

	// ResourceType ResourceType `json:"resource_type" bson:"resource_type,omitempty"`
}
