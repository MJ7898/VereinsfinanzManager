package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Result struct {
	Success string `json:"success"`
}

func SendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetId(r *http.Request) (primitive.ObjectID, error)  {
	vars := mux.Vars(r)

	objID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Errorf("cast ist going wrong %v, STRING: %v", err, objID)}
	return objID, nil
}

func Remove(s []primitive.ObjectID, id primitive.ObjectID) []primitive.ObjectID {
	for i := 0; i <= len(s); i++ {
		if id == s[i] {
			s[i] = s[len(s)-1]
			s[len(s)-1] = primitive.ObjectID{}
			return s[:len(s)-1]
		}
	}
	return s
}
