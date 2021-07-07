package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/utils"
	"net/http"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
)

func CreateClub(w http.ResponseWriter, r *http.Request) {
	log.Infof("Entering createClub-Handler")
	club, err := getClub(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Handler: Club: %v", club)
	if err := service.CreateClub(club); err != nil {
		log.Errorf("Error calling service CreateClub: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Leave createClub-Handler")
}

func getClub(r *http.Request) (*model.Club, error) {
	log.Infof("Entering getClub-Handler")
	var club model.Club
	//ToDo: loh http body ad middleware
	err := json.NewDecoder(r.Body).Decode(&club)
	if err != nil {
		log.Errorf("Can't serialize request body to club struct: %v", err)
		return nil, err
	}
	log.Infof("Leaving getClub-Handler")
	return &club, nil
}

func GetClubs(w http.ResponseWriter, _ *http.Request) {
	log.Infof("Entering GetClubs-Handler")
	clubs, err := service.GetClubs()
	if err != nil {
		log.Errorf("Error calling service GetClubs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clubs); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Infof("Leaving GetClub-Handler")
	utils.SendJson(w, clubs)
}

// GetClub -Handler function to get an single department with id/**
func GetClub(w http.ResponseWriter, r *http.Request) {
	log.Infof("Entering GetClub-Handler")
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)Club: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Infof("Leaving GetClub-Handler")
	club, _ := service.GetClub(id)
	utils.SendJson(w, club)
}

func UpdateClub(w http.ResponseWriter, r *http.Request) {
	log.Infof("Entering UpdateClub-Handler")
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	clubModel, err := getClub(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if clubModel == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	club, err := service.UpdateClub(id, clubModel)
	if err != nil {
		log.Errorf("Failure updateing club with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Leaving UpdateClub-Handler")
	utils.SendJson(w, club)
}

func DeleteClub(w http.ResponseWriter, r *http.Request) {
	log.Infof("Entering DeleteClub-Handler")
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	club, err := service.DeleteClub(id)

	if err != nil {
		log.Errorf("Failure updating club with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if club == nil {
		http.Error(w, "404 club not found", http.StatusNotFound)
		return
	}
	log.Infof("Leaving DeleteClub-Handler")
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}
