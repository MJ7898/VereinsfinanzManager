package model

type Club struct {
	ClubName      string `bson:"club_name,omitempty"`
	ClubLeader    string `bson:"club_leader,omitempty"`
	Budget       float64 `bson:"budget,omitempty"`
	Address       string `bson:"address,omitempty"`
	Description   string `bson:"description,omitempty"`
	Departments []string `bson:"departments,omitempty"`
}
