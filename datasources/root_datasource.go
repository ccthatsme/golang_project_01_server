package datasources

import (
	"github.com/golang_project_01_server/services"
)

type Env struct {
	EmployeeService services.EmployeeService
    AuthorizationService services.AuthService
    ProjectService services.ProjectService
}