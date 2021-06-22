package model

type Team struct {
	// TeamID     int64
	NameOfTeam string   `bson:"name_of_team,omitempty"`
	TeamLeader string   `bson:"team_leader,omitempty"`
	TeamBudget float64  `bson:"team_budget,omitempty"`
	Category   Category `bson:"category,omitempty"`
	Players    []string `bson:"players_id,omitempty"`
}
