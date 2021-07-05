package service

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTeam(team *model.Team) (*mongo.InsertOneResult, error) {
	result, error := client.CreateTeamDB(*team)
	log.Printf("Successfully added team %v to DB", result)
	return result, error
}

func GetTeams() ([]model.Team, error) {
	// var teams []model.Department
	teams, err := client.GetTeamsFromDB()
	if err != nil {
		log.Fatalf("No Documents was found after calling GetTeams: %v", err)
	}
	return teams, nil
}

func GetTeam(id primitive.ObjectID) (model.Team, error) {
	//team := new(model.Department)
	team, err := client.GetTeamWithIDFromDB(id)
	if err != nil {
		log.Printf("Document with ID %V not found! LOG: %v", id, err)
	}
	return team, nil
}

func UpdateTeam(id primitive.ObjectID, team *model.Team) (*model.Team, error) {
	newTeam, _ := client.UpdateTeamFromDB(id, team)
	log.Printf("Output of New Department: %v", newTeam)
	return team, nil
}

func DeleteTeam(id primitive.ObjectID) (*model.Team, error) {
	// client := client.GetMongoDBConnection
	deleteTeamResult, err := client.DeleteTeamDB(id)
	if err != nil {
		log.Fatalf("Error %v was thorwn", err)
	}
	log.Printf("Successfully deleted team %v from DB", deleteTeamResult)
	return nil, nil
}
