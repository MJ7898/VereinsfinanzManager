package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/utils"
	"net/http"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	team, err := getTeam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Handler: Team: %v", team)
	if _, err := service.CreateTeam(team); err != nil {
		log.Errorf("Error calling service CreateTeam: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTeam(r *http.Request) (*model.Team, error) {
	var team model.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Errorf("Can't serialize request body to team struct: %v", err)
		return nil, err
	}
	return &team, nil
}

func GetTeams(w http.ResponseWriter, _ *http.Request) {
	teams, err := service.GetTeams()
	if err != nil {
		log.Errorf("Error calling service GetTeams: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Handler: Return the Result of GET: %v", teams)
	utils.SendJson(w, teams)
}

// GetTeam GetTeam-Handler function to get an single department with id/**
func GetTeam(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)Team: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, _ := service.GetTeam(id)
	utils.SendJson(w, team)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	team, err := getTeam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err = service.UpdateTeam(id, team)
	if err != nil {
		log.Errorf("Failure updating team with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err := service.DeleteTeam(id)

	if err != nil {
		log.Errorf("Failure updating team with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.Error(w, "404 team not found", http.StatusNotFound)
		return
	}
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}

func TeamCost (w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cost, err := service.GetResourceCost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, cost)
}
