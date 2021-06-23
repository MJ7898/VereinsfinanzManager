package model

type Department struct {
	SchemaVersion string `bson:"schema_version,omitempty"`
	// DepartmentID     int64 `bson:"department_id,omitempty"`
	NameOfDepartment string  `bson:"name_of_department,omitempty"`
	DepartmentLeader string  `bson:"department_leader,omitempty"`
	DepartmentBudget float64 `bson:"department_budget,omitempty"`
	//Teams            []Team   `bson:"teams_id,omitempty"`
	Teams []string `bson:"teams_id,omitempty"`
}
