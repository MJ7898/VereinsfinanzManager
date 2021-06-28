package service

import "github.com/MJ7898/VereinsfinanzManager/src/myfinanz/model"

func CreateDepartment(department *model.Department) error  {
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
