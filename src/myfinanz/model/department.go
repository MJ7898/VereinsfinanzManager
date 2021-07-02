package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	SchemaVersion    string `json:"schema_version"`
	Key string `json:"key,omitempty"`
	ID     primitive.ObjectID  `bson:"_id,omitempty"`
	NameOfDepartment string `bson:"name_of_department" json:"name_of_department"`
	DepartmentLeader string `bson:"department_leader" json:"department_leader"`
	DepartmentBudget string `bson:"department_budget" json:"department_budget"`
	//Teams            []Team   `bson:"teams_id,omitempty"`
	// Teams []string `bson:"teams_id,omitempty"`
}
