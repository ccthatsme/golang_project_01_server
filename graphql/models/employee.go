package models

type Employee struct {
	Id                string `json:"id"`
	EmployeeNetworkId string `json:"employeeNetworkId"`
	DisplayName       string `json:"displayName"`
	Email             string `json:"email"`
	Department        string `json:"department"`
	HireDate          string `json:"hireDate"`
	JobTitle          string `json:"jobTitle"`
	Manager           string `json:"manager"`
	Practice          string `json:"practice"`
}

func NewEmployee(emp *Employee) *Employee {

	return &Employee{
		Id:                emp.Id,
		EmployeeNetworkId: emp.EmployeeNetworkId,
		DisplayName:       emp.DisplayName,
		Email:             emp.Email,
		Department:        emp.Department,
		HireDate:          emp.HireDate,
		JobTitle:          emp.JobTitle,
		Manager:           emp.Manager,
		Practice:          emp.Practice,
	}
}
