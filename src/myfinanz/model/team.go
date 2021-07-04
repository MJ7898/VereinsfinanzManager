package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	// TeamID     int64
	SchemaVersion     string   `json:"schema_version" bson:"schema_version"`
	Key string `json:"key,omitempty"`
	ID     primitive.ObjectID  `bson:"_id,omitempty"`
	NameOfTeam        string   `json:"name_of_team" bson:"name_of_team"`
	TeamLeader        string   `json:"team_leader" bson:"team_leader"`
	TeamBudget        float64  `json:"team_budget" bson:"team_budget"`
	// HumanResources    []primitive.ObjectID `json:"human_resources_id" bson:"human_resources_id"`
	// NonHumanResources []primitive.ObjectID `json:"non_human_resources_id" bson:"non_human_resources_id"`
	SumCosts          float64  `json:"overall_costs" bson:"overall_costs"`
}
