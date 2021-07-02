package service

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateDepartment(department *model.Department) error  {
	result := client.CreateDepartmentDB(*department)
	log.Printf("Successfully added department %v to DB", result)
	return result
}

func GetDepartments() ([]model.Department, error)  {
	// var departments []model.Department
	departments, err := client.GetDepartmentsFromDB()
	if err != nil {
		log.Fatalf("No Documents was found after calling GetDepartments: %v", err)
	}
	return departments, nil
}

func GetDepartment(id primitive.ObjectID)(model.Department, error){
	//department := new(model.Department)
	department, err := client.GetDepratmentWithIDFromDB(id)
	if err != nil {
		log.Printf("Document with ID %V not found! LOG: %v", id, err)
	}
	return department, nil
}

func UpdateDepartment(id primitive.ObjectID, department *model.Department) (*model.Department, error)  {
	return department, nil
}

func DeleteDepartment(id primitive.ObjectID) (*model.Department, error)  {
	/*department, err := GetDepartment(id)

	if err == nil {
		return department, nil
	}
	// client := client.GetMongoDBConnection
	deleteDepartmentResult :=  client.DeleteDepartmentDB(*department) // mongoDB.DB.Delete(&department)
	log.Printf("Successfully deleted department %v from DB", deleteDepartmentResult)*/
	return nil, nil
}
