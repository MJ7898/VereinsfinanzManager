package service

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHR(hr *model.HumanResources) error  {
	result := client.CreateHRDB(*hr)
	log.Printf("Successfully added hr %v to DB", result)
	return result
}

func GetHRS() ([]model.HumanResources, error)  {
	// var teams []model.Department
	hrs, err := client.GetHRsFromDB()
	if err != nil {
		log.Fatalf("No Documents was found after calling GetTeams: %v", err)
	}
	return hrs, nil
}

func GetHR(id primitive.ObjectID)(model.HumanResources, error){
	hr, err := client.GetHRWithIDFromDB(id)
	if err != nil {
		log.Printf("Document with ID %V not found! LOG: %v", id, err)
	}
	return hr, nil
}

func DeleteHR(id primitive.ObjectID) (*model.HumanResources, error)  {
	// client := client.GetMongoDBConnection
	deleteHRResult, err :=  client.DeleteHRDB(id)
	if err != nil {
		log.Fatalf("Error %v was thorwn", err)
	}
	log.Printf("Successfully deleted hr %v from DB", deleteHRResult)
	return nil, nil
}

