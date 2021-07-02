package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type result struct {
	Success string `json:"success"`
}

func sendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getId(r *http.Request) (primitive.ObjectID, error)  {
	vars := mux.Vars(r)

	objID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Errorf("cast ist going wrong %v, STRING: %v", err)}
	return objID, nil
}
