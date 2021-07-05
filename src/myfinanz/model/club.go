package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Club struct {
	SchemaVersion string             `bson:"schema_version,omitempty" json:"schema_version"`
	Key           string             `json:"key,omitempty"`
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ClubName      string             `bson:"club_name,omitempty" json:"club_name"`
	ClubLeader    string             `bson:"club_leader,omitempty" json:"club_leader"`
	Budget        float64            `bson:"budget,omitempty" json:"budget"`
	Address       string             `bson:"address,omitempty" json:"address"`
	Description   string             `bson:"description,omitempty" json:"description"`
	Departments   []primitive.ObjectID           `bson:"departments_id,omitempty" json:"departments_id"`
	BankAccount   BankAccount        `bson:"bank_account,omitempty" json:"bank_account"`
}
