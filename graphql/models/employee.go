package models

type Employee struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

func NewEmployee(emp *Employee) *Employee {

	return &Employee{
		Id:          emp.Id,
		DisplayName: emp.DisplayName,
		Email:       emp.Email,
	}
}
