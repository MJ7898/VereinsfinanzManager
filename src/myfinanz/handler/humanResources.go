package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateHR(w http.ResponseWriter, r *http.Request)  {
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
}

func getHR(r *http.Request) (*model.HumanResources, error)  {
	var hr model.HumanResources
	err := json.NewDecoder(r.Body).Decode(&hr)
	if err != nil {
		log.Errorf("Can't serialize request body to HR struct: %v", err)
		return nil, err
	}
	return &hr, nil
}

func GetHRS(w http.ResponseWriter, _ *http.Request)  {
	hr, err := service.GetHRS()
	if err != nil {
		log.Errorf("Error calling service GetHRs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, hr)
}

// GetHR -Handler function to get an single HumanResource with id/**
func GetHR(w http.ResponseWriter, r *http.Request)  {
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)HR: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, _ := service.GetHR(id)
	utils.SendJson(w, team)
}

func DeleteHR(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
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
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}

func AddHR(w http.ResponseWriter, r *http.Request)  {
	teamId, err := utils.GetId(r)
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

	err = service.AddHRDon(teamId, hrDon)
	if err != nil {
		log.Errorf("Failure adding hr to team with ID %v: %v", hrDon, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(w, hrDon)
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
