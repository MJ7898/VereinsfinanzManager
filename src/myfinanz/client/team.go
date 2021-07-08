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

func CreateTeamDB(team model.Team) (*mongo.InsertOneResult, error) {
	client, err := mongoDB.GetMongoClient()
	log.Infof("Error during getMongoClient: %v was thrown", err)
	res, err := client.Database("VfM").Collection("Team").InsertOne(context.TODO(), team) //bson.M{"schema_version":&department.SchemaVersion, "name_of_department": &department.NameOfDepartment, "department_leader": &department.DepartmentLeader, "department_budget": &department.DepartmentBudget})
	if err != nil {
		log.Infof("Error during insertOne: %v was thrown", err)
		log.Infof("Connection isn`t up %v", res)
		return res, err
	}
	entry := log.WithField("ID", team)
	entry.Infof("Successfully added team: %v", res)
	log.Printf("Successfully added Team")
	return res, err
}

// GetTeamWithIDFromDB GetDepratmentWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetTeamWithIDFromDB(id primitive.ObjectID) (model.Team, error) {
	result := model.Team{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database("VfM").Collection("Team")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetTeamsFromDB() ([]model.Team, error) {
	var getResult []model.Team
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("Team")

	cur, errCon := collection.Find(context.TODO(), bson.M{})
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.Team{}
		err := cur.Decode(&jsonRes)
		if err != nil {
			return getResult, err
		}
		getResult = append(getResult, jsonRes)
	}
	log.Printf(" CLient: Return the Result of GET: %v", getResult)
	cur.Close(context.TODO())
	if len(getResult) == 0 {
		return getResult, mongo.ErrNoDocuments
	}
	return getResult, nil
}

func UpdateTeamFromDB(id primitive.ObjectID, team *model.Team) (model.Team, error) {
	result := model.Team{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.M{
		"schema_version":     team.SchemaVersion,
		"name_of_department": team.NameOfTeam,
		"department_leader":  team.TeamLeader,
		"department_budget":  team.TeamBudget,
	}}}
	log.Printf("Result from UPDATER: %v", updater)

	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database("VfM").Collection("Team")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}
	updatedDocu, err := collection.UpdateOne(context.TODO(), filter, updater)
	log.Printf("Updated Document as follow: %v", updatedDocu)

	if err != nil {
		return result, nil
	}
	return model.Team{}, nil
}

func DeleteTeamDB(id primitive.ObjectID) (model.Team, error) {
	result := model.Team{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, err := mongoDB.GetMongoClient()

	res, errCon := client.Database("VfM").
		Collection("Team").
		DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalf("Error: %v was thrown", errCon)
		log.Fatalf("Connection isn`t up %v", res)
		return model.Team{}, nil
	}
	entry := log.WithField("ID", result)
	entry.Infof("Successfully deleted team: %v With ID: %v", result, id)
	log.Printf("Successfully deleted Team")
	return result, nil
}

func TeamCosts(id primitive.ObjectID) ([]bson.M, error){
	filter := bson.M{"team_id": id}
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collectionHR := client.Database("VfM").Collection("HR")

	hrDocument, err := collectionHR.Find(context.TODO(), filter)
	log.Printf("Human Resource %v", hrDocument)
	if err != nil {
		log.Errorf("Error searching HR: %v", err)
		return nil, err
	}

	var hr bson.M
	for hrDocument.Next(context.TODO()) {
		if err = hrDocument.Decode(&hr); err != nil {
			log.Fatal(err)
		}
	}

	hrDocument.Close(context.TODO())

	collectionNHR := client.Database("VfM").Collection("NHR")
	nhrDocument, err := collectionNHR.Find(context.TODO(), filter)

	var nhr bson.M
	for nhrDocument.Next(context.TODO()) {
		if err = nhrDocument.Decode(&nhr); err != nil {
			log.Fatal(err)
		}
	}
	var resourcesFiltered []bson.M

	nhrDocument.Close(context.TODO())

	return append(resourcesFiltered, hr, nhr), nil
}
