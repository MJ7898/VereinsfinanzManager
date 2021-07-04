package client

import (
	"context"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateClubDB(club model.Club) error {
	client, err := GetMongoClient()
	//client, ctx, err := GetMongoDBConnection()
	log.Infof("Client: Error during getMongoClient: %v was thrown", err)
	//log.Infof("Error: %v was thrown", ctx)
	res, err := client.Database("VfM").Collection("Club").InsertOne(context.TODO(), club)
	if err != nil {
		log.Infof("Client: Error during insertOne: %v was thrown", err)
		log.Infof("Client: Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", club)
	entry.Infof("club: %v", club)
	entry.Infof("Successfully added club: %v", res)
	log.Printf("Successfully added Club")
	return err
}

// GetDepratmentWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetClubWithIDFromDB(id primitive.ObjectID) (model.Club, error) {
	result := model.Club{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database("VfM").Collection("Club")
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func GetClubsFromDB() ([]model.Club, error) {
	var getResult []model.Club
	client, err := GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("Club")

	cur, errCon := collection.Find(context.TODO(), bson.M{})
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.Club{}
		err := cur.Decode(&jsonRes)
		if err != nil {
			return getResult, err
		}
		getResult = append(getResult, jsonRes)
	}
	cur.Close(context.TODO())
	if len(getResult) == 0 {
		return getResult, mongo.ErrNoDocuments
	}
	return getResult, nil
}

func UpdateClubFromDB(id primitive.ObjectID, club *model.Club) (model.Club, error) {
	log.Infof("ENTERING UpdateClub-Client")
	result := model.Club{}

	//Adding new Departments via append old slice
	oldDepartments, _ := GetClubWithIDFromDB(id)
	newDepartments := oldDepartments.Departments
	newDepartments = append(newDepartments, club.Departments...)

	bankAccount := model.BankAccount{
		SchemaVersion: club.BankAccount.SchemaVersion,
		BankAccountID: club.BankAccount.BankAccountID,
		OwnerName:     club.BankAccount.OwnerName,
		NameOfBank:    club.BankAccount.NameOfBank,
		Iban:          club.BankAccount.Iban,
	}
	log.Infof("Used BankAccount: %v", bankAccount)
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.M{
		"schema_version": club.SchemaVersion,
		"club_name":      club.ClubName,
		"club_leader":    club.ClubLeader,
		"budget":         club.Budget,
		"address":        club.Address,
		"description":    club.Description,
		"bank_account":   bankAccount,
		//"departments_id": club.Departments,
		"departments_id": newDepartments,
	}}}

	log.Printf("Result from UPDATER: %v", updater)

	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database("VfM").Collection("Club")

	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}
	updatedDocu, err := collection.UpdateOne(context.TODO(), filter, updater)
	log.Printf("Updated Document as follow: %v", updatedDocu)

	if err != nil {
		return result, nil
	}
	//Return result without any error.
	log.Infof("Leaving UpdateClub-Client")
	return model.Club{}, nil
}

func DeleteClubDB(id primitive.ObjectID) (model.Club, error) {
	result := model.Club{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, errCon := GetMongoClient()

	teamsFromDB, err := GetTeamsFromDB()
	if err != nil {
		return model.Club{}, err
	}
	departmentsFromDB, err := GetDepartmentsFromDB()
	if err != nil {
		return model.Club{}, err
	}

	for j := 0; j < len(departmentsFromDB); j++ {
		if departmentsFromDB[j].ID == id {
			if teamsFromDB[j].ID == departmentsFromDB[j].ID {
				deleteAllTeams, err := client.Database("VfM").Collection("Team").DeleteMany(context.TODO(), bson.M{"_id": departmentsFromDB[j].ID})
				if err != nil {
					return model.Club{}, err
				}
				log.Printf("Successfully deleted Collection Department: %v", deleteAllTeams)
			}
			deleteAllDepartments, err := client.Database("VfM").Collection("Department").DeleteMany(context.TODO(), bson.M{"_id":id})
			if err != nil {
				return model.Club{}, err
			}
			log.Printf("Successfully deleted Collection Department: %v", deleteAllDepartments)
		}
	}

	res, errCon := client.Database("VfM").
		Collection("Club").
		DeleteOne(context.TODO(), filter)
	if errCon != nil {
		log.Fatalf("Error: %v was thrown", errCon)
		log.Fatalf("Connection isn`t up %v", res)
		return model.Club{}, nil
	}
	entry := log.WithField("ID", result)
	entry.Infof("Successfully deleted club: %v With ID: %v", result, id)
	log.Printf("Successfully deleted Club")

	return result, nil
}
