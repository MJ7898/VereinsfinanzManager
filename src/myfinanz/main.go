package main

import "github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"

func main() {
	/*log.Println("Starting My-Finanz API Server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}*/

	client.GetMongoDBConnection()

	// collection := client.Database("VfM").Collection("club")
	// res, err := collection.InsertOne(ctx, bson.D{{"club_name", "FC Hackentrick"}, {"club_leader", "Juergen Klopp"}, {"budget", 15000000000}, {"address", "Klosterweg 13, 56789 Heide"}, {"description", "Club Heide for fun and fitness"}, {"departments", bson.A{"Fussball", "Handball", "Tennis", "Golf", "Schwimmen", "Eishockey"}}})
	//id := res.InsertedID

	// fmt.Printf("Successfully inserted %v\n", res)
}
