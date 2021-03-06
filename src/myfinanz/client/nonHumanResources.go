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

func CreateNHRDB(nhr model.NonHumanResources) error {
	client, err := mongoDB.GetMongoClient()
	log.Infof("Error during getMongoClient: %v was thrown", err)
	res, err := client.Database("VfM").Collection("NHR").InsertOne(context.TODO(), nhr)
	if err != nil {
		log.Infof("Error during insertOne: %v was thrown", err)
		log.Infof("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", nhr)
	entry.Infof("Successfully added nhr: %v", res)
	log.Printf("Successfully added NHR")
	return err
}

// GetNHRWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetNHRWithIDFromDB(id primitive.ObjectID) (model.NonHumanResources, error)  {
	result := model.NonHumanResources{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database("VfM").Collection("NHR")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetNHRsFromDB() ([]model.NonHumanResources, error) {
	var getResult []model.NonHumanResources
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("NHR")

	cur, errCon := collection.Find(context.TODO(), bson.M{})
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.NonHumanResources{}
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

func DeleteNHRDB(id primitive.ObjectID,) (model.NonHumanResources, error) {
	result := model.NonHumanResources{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, err := mongoDB.GetMongoClient()

	res, errCon := client.Database("VfM").
		Collection("NHR").
		DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalf("Error: %v was thrown", errCon)
		log.Fatalf("Connection isn`t up %v", res)
		return model.NonHumanResources{}, nil
	}
	entry := log.WithField("ID", result)
	entry.Infof("Successfully deleted NonHumanResource: %v With ID: %v", result, id)
	log.Printf("Successfully deleted NonHumanResource")
	return result, nil
}

func UpdateNHRDBWithTeamDependency(nhr *model.NonHumanResources, teamID primitive.ObjectID) (*model.NonHumanResources, error) {
	nhr.TeamID = teamID

	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return nhr, err
	}
	collection := client.Database("VfM").Collection("NHR")
	createNHR, err := collection.InsertOne(context.TODO(), nhr)
	log.Printf("Successfully insterte Non Human Resourse: %v", createNHR)

	return nhr, nil
}