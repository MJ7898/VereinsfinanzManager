package model

type Team struct {
	// TeamID     int64
	SchemaVersion string  `bson:"schema_version,omitempty"`
	NameOfTeam    string  `bson:"name_of_team,omitempty"`
	TeamLeader    string  `bson:"team_leader,omitempty"`
	TeamBudget    float64 `bson:"team_budget,omitempty"`
	//Category          Category `bson:"category,omitempty"`
	HumanResources    []string `bson:"human_resources_id,omitempty"`
	NonHumanResources []string `bson:"non_human_resources_id,omitempty"`
	SumCosts          float64  `bson:"overall_costs,omitempty"`
	HumanCostsSum     float64  `bson:"human_costs,omitempty"`
	HumanCount        float64  `bson:"human_count,omitempty"`
	//HumanResources    []HumanResources `bson:"human_resources,omitempty"`
	//NonHumanResources []NonHumanResources `bson:"non_human_resources,omitempty"`
}
