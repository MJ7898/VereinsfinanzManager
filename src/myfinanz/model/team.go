package model

type Team struct {
	// TeamID     int64
	NameOfTeam        string   `bson:"name_of_team,omitempty"`
	TeamLeader        string   `bson:"team_leader,omitempty"`
	TeamBudget        float64  `bson:"team_budget,omitempty"`
	Category          Category `bson:"category,omitempty"`
	HumanResources    []string `bson:"human_resources_id,omitempty"`
	NonHumanResources []string `bson:"non_human_resources_id,omitempty"`
	// HumanResources    []HumanResources `bson:"human_resources_id,omitempty"`
	// NonHumanResources []NonHumanResources `bson:"non_human_resources_id,omitempty"`
}
