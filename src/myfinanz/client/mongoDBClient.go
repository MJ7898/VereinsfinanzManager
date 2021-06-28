package client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Client struct {

}

func GetMongoDBConnection(){
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

	err = client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(databases)
	log.Fatalf("Successfully connected to MongoDB %v.", databases)
}