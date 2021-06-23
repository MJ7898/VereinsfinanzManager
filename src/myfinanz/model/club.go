package model

type Club struct {
	SchemaVersion string  `bson:"schema_version,omitempty"`
	ClubName      string  `bson:"club_name,omitempty"`
	ClubLeader    string  `bson:"club_leader,omitempty"`
	Budget        float64 `bson:"budget,omitempty"`
	Address       string  `bson:"address,omitempty"`
	Description   string  `bson:"description,omitempty"`
	//Departments []Department `bson:"departments,omitempty"`
	Departments []string    `bson:"department_id,omitempty"`
	BankAccount BankAccount `bson:"bank_account,omitempty"`
}
