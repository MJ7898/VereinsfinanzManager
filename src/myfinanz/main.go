package main

import (
	"net/http"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	// init database
	defer func() {
		_, err := client.GetMongoClient()
		if err != nil {
			log.Fatalf("Connection to mongoDB refused")
		}
	}()
}

func main() {
	log.Infoln("test: 1, 2, 3 Starting My-Aktion API server")
	router := mux.NewRouter()
	log.Infoln("created new router via mux")
	router.HandleFunc("/health", handler.Health).Methods("GET")

	// clubs
	router.HandleFunc("/club", handler.CreateClub).Methods("POST")
	router.HandleFunc("/clubs", handler.GetClubs).Methods("GET")
	router.HandleFunc("/clubs/{id}", handler.GetClub).Methods("GET")
	router.HandleFunc("/clubs/{id}", handler.UpdateClub).Methods("PUT")
	router.HandleFunc("/clubs/{id}", handler.DeleteClub).Methods("DELETE")

	// departments
	router.HandleFunc("/department", handler.CreateDepartment).Methods("POST")
	router.HandleFunc("/departments", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.GetDepartment).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.UpdateDepartment).Methods("PUT")
	router.HandleFunc("/departments/{id}", handler.DeleteDepartment).Methods("DELETE")

	// teams
	router.HandleFunc("/team", handler.CreateTeam).Methods("POST")
	router.HandleFunc("/teams", handler.GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", handler.GetTeam).Methods("GET")
	router.HandleFunc("/teams/{id}", handler.UpdateTeam).Methods("PUT")
	router.HandleFunc("/teams/{id}", handler.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/teams/{id}/hr", handler.AddHR).Methods("POST")
	router.HandleFunc("/teams/{id}/nhr", handler.AddNHR).Methods("POST")

	// HR
	//router.HandleFunc("/hr", handler.CreateHR).Methods("POST")
	router.HandleFunc("/hrs", handler.GetHRS).Methods("GET")
	router.HandleFunc("/hrs/{id}", handler.GetHR).Methods("GET")
	router.HandleFunc("/hrs/{id}", handler.DeleteHR).Methods("DELETE")

	// NHR
	//router.HandleFunc("/nhr", handler.CreateNHR).Methods("POST")
	router.HandleFunc("/nhrs", handler.GetNHRS).Methods("GET")
	router.HandleFunc("/nhrs/{id}", handler.GetNHR).Methods("GET")
	router.HandleFunc("/nhrs/{id}", handler.DeleteNHR).Methods("DELETE")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
	log.Infof("Server runs on %v", router)
}
