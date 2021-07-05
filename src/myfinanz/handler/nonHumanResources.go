package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateNHR(w http.ResponseWriter, r *http.Request)  {
	log.Infof("Entering createNHR-Handler")
	nhr, err := getNHR(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Handler: NHR: %v", nhr)
	if err := service.CreateNHR(nhr); err != nil {
		log.Errorf("Error calling service CreateNHR: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Leave createNHR-Handler")
}

func getNHR(r *http.Request) (*model.NonHumanResources, error)  {
	log.Infof("Entering getNHR-Handler")
	var nhr model.NonHumanResources
	//ToDo: loh http body ad middleware
	err := json.NewDecoder(r.Body).Decode(&nhr)
	if err != nil {
		log.Errorf("Can't serialize request body to nhr struct: %v", err)
		return nil, err
	}
	log.Infof("Leaving getNHR-Handler")
	return &nhr, nil
}

func GetNHRS(w http.ResponseWriter, _ *http.Request)  {
	nhr, err := service.GetNHRS()
	if err != nil {
		log.Errorf("Error calling service GetNHRs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json. NewEncoder(w).Encode(nhr); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sendJson(w, nhr)
}

// GetHR -Handler function to get an single HumanResource with id/**
func GetNHR(w http.ResponseWriter, r *http.Request)  {
	id, err := getId(r)
	// var objectResult primitive.ObjectID = id
	if err != nil {
		log.Errorf("Error calling servie Get(Single)NHR: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team, _ := service.GetNHR(id)
	sendJson(w, team)
}

func DeleteNHR(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nhr, err := service.DeleteNHR(id)

	if err != nil {
		log.Errorf("Failure updating hr with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if nhr == nil {
		http.Error(w, "404 nhr not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "Success (Ok)"})
}
