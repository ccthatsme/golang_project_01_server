package services

import (
	"github.com/golang_project_01_server/global_methods"
	"github.com/golang_project_01_server/graphql/models"
	"encoding/json"
	"fmt"
)

type ProjectService interface {
	GetAllProjects(token string) []models.Project
	GetProject(token string, id string) models.Project
}

type ProjectDatasource struct {
	CompanyHttp global_methods.HttpDataSource
}

func NewProService(companyHttp global_methods.HttpDataSource) ProjectService {
	return &ProjectDatasource{
		CompanyHttp: companyHttp,
	}
}

var proList []models.Project
var pro models.Project

func (ds *ProjectDatasource) GetAllProjects(token string) []models.Project {

	result, err := ds.CompanyHttp.Get("/employees", token)

	json.Unmarshal(result, &proList)
	if err != nil {
		fmt.Println("line 58")
	}
	return proList

}

func (ds *ProjectDatasource) GetProject(token string, id string) models.Project {

	result, err := ds.CompanyHttp.Get("/employees/"+id, token)

	json.Unmarshal(result, &pro)
	if err != nil {
		fmt.Println("line 58")
	}
	return pro

}