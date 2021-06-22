package model

type Department struct {
	// DepartmentID     int64 `bson:"department_id,omitempty"`
	NameOfDepartment string   `bson:"name_of_department,omitempty"`
	DepartmentLeader string   `bson:"department_leader,omitempty"`
	DepartmentBudget float64  `bson:"department_budget,omitempty"`
	Category         Category `bson:"category,omitempty"`
	Teams            []Team   `bson:"teams_id,omitempty"`
	// Teams            []string `bson:"teams_id,omitempty"`
}