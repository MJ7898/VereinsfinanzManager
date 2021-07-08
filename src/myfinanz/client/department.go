package client

import (
	"context"
	"sync"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/mongoDB"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

func CreateDepartmentDB(department model.Department) error {
	client, err := mongoDB.GetMongoClient()
	//client, ctx, err := GetMongoDBConnection()
	log.Infof("Error during getMongoClient: %v was thrown", err)
	//log.Infof("Error: %v was thrown", ctx)
	res, err := client.Database("VfM").Collection("Department").InsertOne(context.TODO(), department) //bson.M{"schema_version":&department.SchemaVersion, "name_of_department": &department.NameOfDepartment, "department_leader": &department.DepartmentLeader, "department_budget": &department.DepartmentBudget})
	if err != nil {
		log.Infof("Error during insertOne: %v was thrown", err)
		log.Infof("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", department)
	entry.Infof("Successfully added department: %v", res)
	log.Printf("Successfully added Department")
	return err
}

// GetDepratmentWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetDepratmentWithIDFromDB(id primitive.ObjectID) (model.Department, error) {
	result := model.Department{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database("VfM").Collection("Department")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetDepartmentsFromDB() ([]model.Department, error) {
	var getResult []model.Department
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("Department")

	cur, errCon := collection.Find(context.TODO(), bson.M{})
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.Department{}
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

func UpdateDepartmentFromDB(id primitive.ObjectID, department *model.Department) (model.Department, error) {
	result := model.Department{}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.M{
		"schema_version":     department.SchemaVersion,
		"name_of_department": department.NameOfDepartment,
		"department_leader":  department.DepartmentLeader,
		"department_budget":  department.DepartmentBudget,
		"teams_id":           department.Teams,
		"department_cost":    department.DepartmentCost,
	}}}
	log.Printf("Result from UPDATER: %v", updater)

	//Get MongoDB connection using connectionhelper.
	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database("VfM").Collection("Department")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}
	updatedDocu, err := collection.UpdateOne(context.TODO(), filter, updater)
	log.Printf("Updated Document as follow: %v", updatedDocu)

	if err != nil {
		return result, nil
	}
	return model.Department{}, nil
}

func DeleteDepartmentDB(id primitive.ObjectID) (model.Department, error) {
	result := model.Department{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, errCon := mongoDB.GetMongoClient()

	res, errCon := client.Database("VfM").
		Collection("Department").
		DeleteOne(context.TODO(), filter)
	if errCon != nil {
		log.Fatalf("Error: %v was thrown", errCon)
		log.Fatalf("Connection isn`t up %v", res)
		return model.Department{}, nil
	}
	entry := log.WithField("ID", result)
	entry.Infof("Successfully deleted department: %v With ID: %v", result, id)
	log.Infof("Successfully deleted Department")
	return result, nil
}

func UpdateCosts(insertedTeam *mongo.InsertOneResult, team *model.Team, departmentID primitive.ObjectID) error {
	result, err := GetDepratmentWithIDFromDB(departmentID)
	if err != nil {
		return err
	}

	teamID, _ := insertedTeam.InsertedID.(primitive.ObjectID)
	log.Printf("Team ID : %v", teamID)
	result.Teams = append(result.Teams, teamID)
	log.Printf("Team ID Array: %v", result.Teams)

	log.Printf("Team Costs old: %v", result.DepartmentCost)
	result.DepartmentCost += team.SumCosts
	log.Printf("Team Costs new: %v", result.DepartmentCost)

	updater := bson.D{primitive.E{Key: "$set", Value: bson.M{
		"schema_version":     result.SchemaVersion,
		"name_of_department": result.NameOfDepartment,
		"department_leader":  result.DepartmentLeader,
		"department_budget":  result.DepartmentBudget,
		"teams_id":           result.Teams,
		"department_cost":    result.DepartmentCost,
	}}}

	log.Printf("Result from UPDATER: %v", updater)

	client, err := mongoDB.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database("VfM").Collection("Department")
	filterClub := bson.D{primitive.E{Key: "_id", Value: result.ID}}
	updatedClub, err := collection.UpdateOne(context.TODO(), filterClub, updater)
	log.Printf("Updated Document as follow: %v", updatedClub)

	if err != nil {
		return err
	}
	return nil
}
