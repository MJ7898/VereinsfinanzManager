package service

import (
	"log"

	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDepartment(department *model.Department) error {
	result := client.CreateDepartmentDB(*department)
	log.Printf("Successfully added department %v to DB", result)
	return result
}

func GetDepartments() ([]model.Department, error) {
	departments, err := client.GetDepartmentsFromDB()
	if err != nil {
		log.Printf("No Documents was found after calling GetDepartments: %v", err)
	}
	return departments, nil
}

func GetDepartment(id primitive.ObjectID) (model.Department, error) {
	department, err := client.GetDepratmentWithIDFromDB(id)
	if err != nil {
		log.Printf("Document with ID %v not found! LOG: %v", id, err)
	}
	return department, nil
}

func UpdateDepartment(id primitive.ObjectID, department *model.Department) (*model.Department, error) {
	newDepartment, _ := client.UpdateDepartmentFromDB(id, department)
	log.Printf("Output of New Department: %v", newDepartment)
	return department, nil
}

func DeleteDepartment(id primitive.ObjectID) (*model.Department, error) {
	deleteDepartmentResult, err := client.DeleteDepartmentDB(id)
	if err != nil {
		log.Fatalf("Error %v was thorwn", err)
	}
	log.Printf("Successfully deleted department %v from DB", deleteDepartmentResult)
	updateClub, err := client.UpdateClubFromDBRemove(id)
	log.Printf("successfully deleted Department and updated Club: %v", updateClub)
	return nil, nil
}

func AddTeamWithDepartment(departmentID primitive.ObjectID, team *model.Team) error {
	result, error := client.CreateTeamDB(*team)
	error = client.UpdateCosts(result, team, departmentID)
	return error
}
