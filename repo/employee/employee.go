package employee

import (
	"fmt"
	"github.com/golang_project_01_server/Models"
)

type EmployeeService interface {
	GetAllEmployees() ([]*Models.Employee, err)
}

type EmployeeDatasource struct {
	CompanyHttp *global.HttpDataSource
}

func NewEmployeeService(companyHttp *global.HttpDataSource) EmployeeService {
	return &EmployeeDatasource{
		CompanyHttp: companyHttp,
	}
}

var allEmp []Models.Employee

func (e *EmployeeDatasource) GetAllEmployees(w http.ResponseWriter, r *http.Request) ([]*Models.Employee, err) {
	data, err := e.CompanyHttp.Get(("GetAllEmployees"))
	if err != nil {
		fmt.Println("employee repo")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")

	ctx := r.Context()
	authorization := ctx.Value("X-Authorization")
	t := authorization.(string)

	// 	array := datasources.GetAllEmployees(t)

	json.Unmarshal(data, &allEmp)
	if err != nil {
		fmt.Println("line 58")
	}
	return allEmp

	//json.NewEncoder(w).Encode(array)

}
