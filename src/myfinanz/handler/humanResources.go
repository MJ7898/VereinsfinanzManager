package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateHR(w http.ResponseWriter, r *http.Request)  {
	log.Infof("Entering createHR-Handler")
	hr, err := getHR(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Handler: HR: %v", hr)
	if err := service.CreateHR(hr); err != nil {
		log.Errorf("Error calling service CreateHR: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Leave createHR-Handler")
}

func getHR(r *http.Request) (*model.HumanResources, error)  {
	log.Infof("Entering getHR-Handler")
	var hr model.HumanResources
	//ToDo: loh http body ad middleware
	err := json.NewDecoder(r.Body).Decode(&hr)
	if err != nil {
		log.Errorf("Can't serialize request body to HR struct: %v", err)
		return nil, err
	}
	log.Infof("Leaving getHR-Handler")
	return &hr, nil
}

func GetHRS(w http.ResponseWriter, _ *http.Request)  {
	hr, err := service.GetHRS()
	if err != nil {
		log.Errorf("Error calling service GetHRs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json. NewEncoder(w).Encode(hr); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sendJson(w, hr)
}

// GetHR -Handler function to get an single HumanResource with id/**
func GetHR(w http.ResponseWriter, r *http.Request)  {
	id, err := getId(r)
	// var objectResult primitive.ObjectID = id
	if err != nil {
		log.Errorf("Error calling servie Get(Single)HR: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team, _ := service.GetHR(id)
	sendJson(w, team)
}

func DeleteHR(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hr, err := service.DeleteHR(id)

	if err != nil {
		log.Errorf("Failure updating hr with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if hr == nil {
		http.Error(w, "404 hr not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "Success (Ok)"})
}

func AddHR(w http.ResponseWriter, r *http.Request)  {
	teamId, err := getId(r)
	if err != nil {
		log.Errorf("Error getting ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hrDon, err := getHRDon(r)
	if err != nil {
		log. Errorf("Can't serialize body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: if the hrDon doesn't exist, return 404 - don't show FK error
	err = service.AddHRDon(teamId, hrDon)
	if err != nil {
		log.Errorf("Failure adding hr to team with ID %v: %v", hrDon, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, hrDon)
}

func getHRDon(r *http.Request) (*model.HumanResources, error) {
	var hr model.HumanResources
	err := json.NewDecoder(r.Body).Decode(&hr)
	if err != nil {
		log.Errorf("Can't serialize request body to hr struct: %v", err)
		return nil, err
	}
	return &hr, nil
}
