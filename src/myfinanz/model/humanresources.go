package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HumanResources struct {
	SchemaVersion    string             `json:"schema_version" bson:"schema_version,omitempty"`
	Key              string             `json:"key,omitempty"`
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"player_name,omitempty"`
	Value            float64            `json:"value" bson:"value,omitempty"`
	Salary           float64            `json:"salary" bson:"salary,omitempty"`
	ContractRuntime  string             `json:"contract_runtime" bson:"contract_runtime,omitempty"`
	TeamID           primitive.ObjectID `json:"team_id" bson:"team_id"`
}
