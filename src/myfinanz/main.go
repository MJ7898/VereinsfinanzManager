package main

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting My-Finanz API Server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
