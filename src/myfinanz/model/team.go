package model

type Team struct {
	TeamID     int64
	NameOfTeam string
	TeamLeader string
	TeamBudget float64
	Category   Category
}
