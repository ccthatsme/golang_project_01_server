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


