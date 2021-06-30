package service

import (
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/client"
	"github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"
	"log"
)

func CreateDepartment(department *model.Department) error  {
	result := client.CreateDepartmentDB(*department)
	log.Printf("Successfully added department %v to DB", result)
	return nil
}

func GetDepartments() ([]model.Department, error)  {
	var departments []model.Department
	return departments, nil
}

func GetDepartment(id uint)(*model.Department, error){
	department := new(model.Department)
	return department, nil
}

func UpdateDepartment(id uint, department *model.Department) (*model.Department, error)  {
	return department, nil
}

func DeleteDepartment(id uint) (*model.Department, error)  {
	department, err := GetDepartment(id)

	if err == nil {
		return department, nil
	}
	// client := client.GetMongoDBConnection
	deleteDepartmentResult :=  client.DeleteDepartmentDB(*department) // mongoDB.DB.Delete(&department)
	log.Printf("Successfully deleted department %v from DB", deleteDepartmentResult)
	return department, nil
}
