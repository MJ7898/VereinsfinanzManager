package client

import (
	"context"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/mongoDB"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateHRDB(hr model.HumanResources) error {
	client, err := mongoDB.GetMongoClient()
	//client, ctx, err := GetMongoDBConnection()
	log.Infof("Error during getMongoClient: %v was thrown", err)
	//log.Infof("Error: %v was thrown", ctx)
	res, err := client.Database("VfM").Collection("HR").InsertOne(context.TODO(), hr)
	if err != nil {
		log.Infof("Error during insertOne: %v was thrown", err)
		log.Infof("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", hr)
	entry.Infof("Successfully added team: %v", res)
	log.Printf("Successfully added Team")
	return err
}

// GetHRWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetHRWithIDFromDB(id primitive.ObjectID) (model.HumanResources, error)  {
	result := model.HumanResources{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Get MongoDB connection using connectionhelper.
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database("VfM").Collection("HR")
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func GetHRsFromDB() ([]model.HumanResources, error) {
	var getResult []model.HumanResources
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("HR")

	cur, errCon := collection.Find(context.TODO(), bson.M{})
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.HumanResources{}
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

func DeleteHRDB(id primitive.ObjectID,) (model.HumanResources, error) {
	result := model.HumanResources{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, err := mongoDB.GetMongoClient()

	res, errCon := client.Database("VfM").
		Collection("HR").
		DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalf("Error: %v was thrown", errCon)
		log.Fatalf("Connection isn`t up %v", res)
		return model.HumanResources{}, nil
	}
	entry := log.WithField("ID", result)
	entry.Infof("Successfully deleted HumanResource: %v With ID: %v", result, id)
	log.Printf("Successfully deleted HumanResource")
	return result, nil
}

func UpdateHRDBWithTeamDependency(hr *model.HumanResources, teamID primitive.ObjectID) (*model.HumanResources, error) {
	hr.TeamID = teamID

	//Get MongoDB connection using connectionhelper.
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return hr, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database("VfM").Collection("HR")
	//Perform FindOne operation & validate against the error.
	createHR, err := collection.InsertOne(context.TODO(), hr)
	log.Printf("Successfully insterte Human Resourse: %v", createHR)

	return hr, nil
}
