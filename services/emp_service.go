package services

import (
	"github.com/golang_project_01_server/global_methods"
	"github.com/golang_project_01_server/graphql/models"
	"encoding/json"
	"fmt"
)

type EmployeeService interface {
	GetAllEmployees(token string) []models.Employee
	GetEmployee(token string, empNetworkId string) models.Employee
}

type EmployeeDataSource struct {
	CompanyHttp global_methods.HttpDataSource
}

func NewEmpService(companyHttp global_methods.HttpDataSource) EmployeeService {
	return &EmployeeDataSource{
		CompanyHttp: companyHttp,
	}
}

var emp []models.Employee
var singleEmp models.Employee

func (ds *EmployeeDataSource) GetAllEmployees(token string) []models.Employee {

	result, err := ds.CompanyHttp.Get("/employees", token)

	json.Unmarshal(result, &emp)
	if err != nil {
		fmt.Println("line 58")
	}
	return emp

}

func (ds *EmployeeDataSource) GetEmployee(token string, empNetworkId string) models.Employee {

	result, err := ds.CompanyHttp.Get("/employees/"+empNetworkId, token)

	json.Unmarshal(result, &singleEmp)
	if err != nil {
		fmt.Println("line 58")
	}
	return singleEmp

}