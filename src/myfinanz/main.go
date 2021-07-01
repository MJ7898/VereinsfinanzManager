package main

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	// init database
	defer func() {
		_, _, err := client.GetMongoDBConnection()
		if err != nil {
			log.Fatalf("Connection to mongoDB refused")
		}
	}()
	// init logger
	/*log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)*/
}
func main() {
	log.Infoln("test: 1, 2, 3 Starting My-Aktion API server")
	router := mux.NewRouter()
	log.Infoln("created new router via mux")
	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/department", handler.CreateDepartment).Methods("POST")
	router.HandleFunc("/departments", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.GetDepartment).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.UpdateDepartment).Methods("PUT")
	router.HandleFunc("/departments/{id}", handler.DeleteDepartment).Methods("DELETE")
	// router.HandleFunc("/campaigns/{id}/donation", handler.AddDonation).Methods("POST")
	// go monitortransactions()
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}

	log.Infof("Server runs on %v", router)

	// client.GetMongoDBConnection()

	// collection := client.Database("VfM").Collection("club")
	// res, err := collection.InsertOne(ctx, bson.D{{"club_name", "FC Hackentrick"}, {"club_leader", "Juergen Klopp"}, {"budget", 15000000000}, {"address", "Klosterweg 13, 56789 Heide"}, {"description", "Club Heide for fun and fitness"}, {"departments", bson.A{"Fussball", "Handball", "Tennis", "Golf", "Schwimmen", "Eishockey"}}})
	//id := res.InsertedID

	// fmt.Printf("Successfully inserted %v\n", res)
}

