package service

import (
	"log"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateClub(club *model.Club) error {
	result := client.CreateClubDB(*club)
	log.Printf("Service: Successfully added club %v to DB", result)
	return result
}

func GetClubs() ([]model.Club, error) {
	// var departments []model.Department
	clubs, err := client.GetClubsFromDB()
	if err != nil {
		log.Fatalf("No Documents was found after calling GetClubs-Service: %v", err)
	}
	return clubs, nil
}

func GetClub(id primitive.ObjectID) (model.Club, error) {
	//department := new(model.Department)
	club, err := client.GetClubWithIDFromDB(id)
	if err != nil {
		log.Printf("Service: Document with ID %v not found! LOG: %v", id, err)
	}
	return club, nil
}

func UpdateClub(id primitive.ObjectID, club *model.Club) (model.Club, error) {
	newClub, _ := client.UpdateClubFromDB(id, club)
	log.Printf("Service: Output of New Department: %v", newClub)
	return newClub, nil
}

func DeleteClub(id primitive.ObjectID) (*model.Club, error) {
	// client := client.GetMongoDBConnection
	deleteClubResult, err := client.DeleteClubDB(id)
	if err != nil {
		log.Fatalf("Error %v was thorwn", err)
	}
	log.Printf("Service: Successfully deleted department %v from DB", deleteClubResult)
	return nil, nil
}
