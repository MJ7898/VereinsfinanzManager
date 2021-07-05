package model

type Role string
// TODO: Use this model in combination with the account

const (
	DEPARTMENTORGANIZER Role = "Department-Organizer"
	TEAMORGANIZER       Role = "Team-Organizer"
	MEMBER               Role = "Member"
)
