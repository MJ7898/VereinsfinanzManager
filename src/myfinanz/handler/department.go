package handler

import (
	"encoding/json"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateDepartment(w http.ResponseWriter, r *http.Request)  {
	department, err := getDepartment(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Infof("Department: %v", department)
	if err := service.CreateDepartment(department); err != nil {
		log.Errorf("Error calling service CreateDepartment: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getDepartment(r *http.Request) (*model.Department, error)  {
	var department model.Department
	//ToDo: loh http body ad middleware
	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		log.Errorf("Can't serialize request body to department struct: %v", err)
		return nil, err
	}
	return &department, nil
}

func GetDepartments(w http.ResponseWriter, _ *http.Request)  {
	departments, err := service.GetDepartments()
	if err != nil {
		log.Errorf("Error calling service GetDepartments: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json. NewEncoder(w).Encode(departments); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sendJson(w, departments)
}

// GetDepartment GetDepartment-Handler function to get an single department with id/**
func GetDepartment(w http.ResponseWriter, r *http.Request)  {
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error calling servie Get(Single)Department: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	department, err := service.GetDepartment(id)
	if department == nil {
		http.Error(w,"404 Department not found", http.StatusNotFound)
		return
	}
	sendJson(w, department)
}

func UpdateDepartment(w http.ResponseWriter, r *http.Request)  {
	id, err := getId(r)
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
	sendJson(w, department)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
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
	sendJson(w, result{Success: "Success (Ok)"})
}

//ToDo: Add Update and Delete Department...
