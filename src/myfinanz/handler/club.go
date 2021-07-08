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
}

func getClub(r *http.Request) (*model.Club, error) {
	var club model.Club
	err := json.NewDecoder(r.Body).Decode(&club)
	if err != nil {
		log.Errorf("Can't serialize request body to club struct: %v", err)
		return nil, err
	}
	return &club, nil
}

func GetClubs(w http.ResponseWriter, _ *http.Request) {
	clubs, err := service.GetClubs()
	if err != nil {
		log.Errorf("Error calling service GetClubs: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, clubs)
}

// GetClub -Handler function to get an single department with id/**
func GetClub(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)Club: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	club, _ := service.GetClub(id)
	utils.SendJson(w, club)
}

func UpdateClub(w http.ResponseWriter, r *http.Request) {
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
	utils.SendJson(w, club)
}

func DeleteClub(w http.ResponseWriter, r *http.Request) {
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
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}
