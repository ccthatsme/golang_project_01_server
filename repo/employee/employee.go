package employee

import (
"github.com/golang_project_01_server/Models"

)

type EmployeeService interface {
GetAllEmployees() ([]*Models.Employee, err)

}


type EmployeeDatasource struct {
	CompanyHttp *global.HttpDataSource
}

func (e *EmployeeDatasource) GetAllEmployees() ([]*Models.Employee, err){

}