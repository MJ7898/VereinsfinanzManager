package client

import (
	"context"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Client struct {

}

func GetMongoDBConnection() (mongo.Client,context.Context, error) {
	// "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))//"mongodb+srv://VfMAdmin:VfMAdmin@cluster-vfm.xlagy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)
	return *client, ctx, nil
	/*err = client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(databases)
	log.Fatalf("Successfully connected to MongoDB %v.", databases)*/
}
func CreateDepartmentDB(department model.Department) error {
	client, ctx, err := GetMongoDBConnection()

	res, err := client.Database("VfM").Collection("Department").InsertOne(ctx, bson.A{department})
	if err != nil {
		log.Fatalf("Error: %v was thrown", err)
		log.Fatalf("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", department)
	entry.Infof("Successfully added department: %v", res)
	log.Printf("Successfully added Department")
	return err
}

func UpdateDepartmentDB()  {

}


func DeleteDepartmentDB(department model.Department) error {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, ctx, error := GetMongoDBConnection()

	res, err := client.Database("VfM").Collection("Department").DeleteOne(ctx, bson.M{"name_of_department": department.NameOfDepartment})
	if err != nil {
		log.Fatalf("Error: %v was thrown", error)
		log.Fatalf("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", department)
	entry.Infof("Successfully deleted department: %v", department.NameOfDepartment)
	log.Printf("Successfully deleted Department")
	return err
}