package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/utils"
	"net/http"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
)

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	department, err := getDepartment(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Handler: Department: %v", department)
	if err := service.CreateDepartment(department); err != nil {
		log.Errorf("Error calling service CreateDepartment: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getDepartment(r *http.Request) (*model.Department, error) {
	var department model.Department
	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		log.Errorf("Can't serialize request body to department struct: %v", err)
		return nil, err
	}
	return &department, nil
}

func GetDepartments(w http.ResponseWriter, _ *http.Request) {
	departments, err := service.GetDepartments()
	if err != nil {
		log.Errorf("Error calling service GetDepartments: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, departments)
}

// GetDepartment GetDepartment-Handler function to get an single department with id/**
func GetDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)Department: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	department, _ := service.GetDepartment(id)
	utils.SendJson(w, department)
}

func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	department, err := getDepartment(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	department, err = service.UpdateDepartment(id, department)
	if err != nil {
		log.Errorf("Failure updateing department with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if department == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendJson(w, department)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	department, err := service.DeleteDepartment(id)

	if err != nil {
		log.Errorf("Failure updating department with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if department == nil {
		http.Error(w, "404 department not found", http.StatusNotFound)
		return
	}
	utils.SendJson(w, utils.Result{Success: "Success (Ok)"})
}

func AddTeamWithDepartment(w http.ResponseWriter, r *http.Request) {
	departmentID, err := utils.GetId(r)
	if err != nil {
		log.Errorf("Error getting ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err := getTeam(r)
	if err != nil {
		log.Errorf("Can't serialize body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.AddTeamWithDepartment(departmentID, team)
	if err != nil {
		log.Errorf("Handler: Add was not successfull")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
