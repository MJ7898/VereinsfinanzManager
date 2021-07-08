package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateNHR(w http.ResponseWriter, r *http.Request)  {
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
}

func getNHR(r *http.Request) (*model.NonHumanResources, error)  {
	var nhr model.NonHumanResources
	err := json.NewDecoder(r.Body).Decode(&nhr)
	if err != nil {
		log.Errorf("Can't serialize request body to nhr struct: %v", err)
		return nil, err
	}
	return &nhr, nil
}

func GetNHRS(w http.ResponseWriter, _ *http.Request)  {
	nhr, err := service.GetNHRS()
	if err != nil {
		log.Errorf("Error calling service GetNHRs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, nhr)
}

// GetNHR -Handler function to get an single HumanResource with id/**
func GetNHR(w http.ResponseWriter, r *http.Request)  {
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)NHR: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team, _ := service.GetNHR(id)
	utils.SendJson(w, team)
}

func DeleteNHR(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
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
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}

func AddNHR(w http.ResponseWriter, r *http.Request)  {
	teamId, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error getting ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nhrDon, err := getNHRDon(r)
	if err != nil {
		log. Errorf("Can't serialize body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AddNHRDon(teamId, nhrDon)
	if err != nil {
		log.Errorf("Failure adding nhr to team with ID %v: %v", nhrDon, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(w, nhrDon)
}

func getNHRDon(r *http.Request) (*model.NonHumanResources, error) {
	var nhr model.NonHumanResources
	err := json.NewDecoder(r.Body).Decode(&nhr)
	if err != nil {
		log.Errorf("Can't serialize request body to nhr struct: %v", err)
		return nil, err
	}
	return &nhr, nil
}

