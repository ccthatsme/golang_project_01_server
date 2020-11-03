package repo

import (
	"github.com/golang_project_01_server/repo/employee"
)

type Resolver struct {
	EmployeeService employee.EmployeeService
}
