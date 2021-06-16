package model

type Department struct {
	DepartmentID     int64
	NameOfDepartment string
	DepartmentLeader string
	DepartmentBudget float64
	Category         Category
}
