package service

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNHR(nhr *model.NonHumanResources) error  {
	result := client.CreateNHRDB(*nhr)
	log.Printf("Successfully added hr %v to DB", result)
	return result
}

func GetNHRS() ([]model.NonHumanResources, error)  {
	// var teams []model.Department
	nhrs, err := client.GetNHRsFromDB()
	if err != nil {
		log.Fatalf("No Documents was found after calling GetNHRSs: %v", err)
	}
	return nhrs, nil
}

func GetNHR(id primitive.ObjectID)(model.NonHumanResources, error){
	nhr, err := client.GetNHRWithIDFromDB(id)
	if err != nil {
		log.Printf("Document with ID %V not found! LOG: %v", id, err)
	}
	return nhr, nil
}

func DeleteNHR(id primitive.ObjectID) (*model.NonHumanResources, error)  {
	// client := client.GetMongoDBConnection
	deleteNHRResult, err :=  client.DeleteNHRDB(id)
	if err != nil {
		log.Fatalf("Error %v was thorwn", err)
	}
	log.Printf("Successfully deleted nhr %v from DB", deleteNHRResult)
	return nil, nil
}

func AddNHRDon(id primitive.ObjectID, nhrDon *model.NonHumanResources) error  {
	team, err := GetTeam(id)
	if err != nil {
		return err
	}

	// team.ID := teamId
	result, _ := client.UpdateNHRDBWithTeamDependency(nhrDon, team.ID)
	log.Printf("Successfully added human ressource and put there the team id: %v", result)
	return nil
}
