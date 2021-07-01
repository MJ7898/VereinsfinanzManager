package model

type Department struct {
	SchemaVersion    string `json:"schema_version"`
	DepartmentID     string  `json:"_id,omitempty"`
	NameOfDepartment string `json:"name_of_department"`
	DepartmentLeader string `json:"department_leader"`
	DepartmentBudget string `json:"department_budget"`
	//Teams            []Team   `bson:"teams_id,omitempty"`
	// Teams []string `bson:"teams_id,omitempty"`
}
