package client

import (
	"context"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)


/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client
//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error
//Used to execute client creation procedure only once.
var mongoOnce sync.Once
//I have used below constants just to hold required database config's.


func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil); if err == nil { log.Infoln("Ping Successful")}
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func GetMongoDBConnection() (mongo.Client,context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { return *client, ctx, err }

	/* "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))//"mongodb+srv://VfMAdmin:VfMAdmin@cluster-vfm.xlagy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}*/
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
	client, err := GetMongoClient()
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

 /*func GetDepratmentWithIDFromDB(id string, objectID string) ([]model.Department, error)  {
	var getResult []model.Department
	client, err := GetMongoClient()
	if err != nil {
		log.Errorf("Cannot Connect to DB: %v", err)
	}
	collection := client.Database("VfM").Collection("Department")

	var departmentFilter []model.Department
	var bSonFilter []byte
	cur, err := collection.Find(context.TODO(), bson.M{"ID":id}json.Unmarshal(bSonFilter, &departmentFilter ))
	if err != nil {
		log.Fatal(err)
	}

	var departmentsFiltered []model.Department
	if err = cur.All(context.TODO(), &departmentsFiltered); err != nil{
		log.Fatal(err)
	}
	log.Printf("Document with _id: %v was found", cur)
	if errCon != nil {
		return getResult, errCon
	}
	for cur.Next(context.TODO()) {
		jsonRes := model.Department{}
		err := cur.Decode(&jsonRes)
		if err != nil {
			return getResult, err
		}
		log.Printf("Cursor found Document with _id: %v and append it to the result Array", cur)
		getResult = append(getResult, jsonRes)
	}
	cur.Close(context.TODO())
	if len(getResult) == 0 {
		log.Printf("Eorror: Cursor is null and has no Content from _id: %v", getResult)
		return getResult, mongo.ErrNoDocuments
	}
	return departmentsFiltered, nil
}*/

// GetDepratmentWithIDFromDB  GetIssuesByCode - Get All issues for collection
func GetDepratmentWithIDFromDB(id primitive.ObjectID) (model.Department, error)  {
		result := model.Department{}
		//Define filter query for fetching specific document from collection
		filter := bson.D{primitive.E{Key: "_id", Value: id}}
		//Get MongoDB connection using connectionhelper.
		client, err := GetMongoClient()
		if err != nil {
			return result, err
		}
		//Create a handle to the respective collection in the database.
		collection := client.Database("VfM").Collection("Department")
		//Perform FindOne operation & validate against the error.
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			return result, err
		}
		//Return result without any error.
		return result, nil
}

func GetDepartmentsFromDB() ([]model.Department, error) {
	var getResult []model.Department
	client, err := GetMongoClient()
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

func UpdateDepartmentDB()  {

}


func DeleteDepartmentDB(department model.Department) error {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	/*client, ctx, error := GetMongoDBConnection()

	res, err := client.Database("VfM").Collection("Department").DeleteOne(ctx, bson.M{"name_of_department": department.NameOfDepartment})
	if err != nil {
		log.Fatalf("Error: %v was thrown", error)
		log.Fatalf("Connection isn`t up %v", res)
		return err
	}
	entry := log.WithField("ID", department)
	entry.Infof("Successfully deleted department: %v", department.NameOfDepartment)
	log.Printf("Successfully deleted Department")*/
	return nil
}