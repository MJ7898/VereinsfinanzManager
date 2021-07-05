package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	SchemaVersion    string             `json:"schema_version"`
	Key              string             `json:"key,omitempty"`
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	NameOfDepartment string             `bson:"name_of_department" json:"name_of_department"`
	DepartmentLeader string             `bson:"department_leader" json:"department_leader"`
	DepartmentBudget float64             `bson:"department_budget" json:"department_budget"`
	DepartmentCost   float64            `json:"department_cost" bson:"department_cost"`
	//Teams            []Team   `bson:"teams_id,omitempty"`
	Teams []primitive.ObjectID  `json:"teams_id" bson:"teams_id,omitempty"`
}
