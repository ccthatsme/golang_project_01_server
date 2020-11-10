package resolver

import (
	"context"
	//"fmt"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

type employeeResolver struct {
	Employee     *models.Employee
	rootResolver *Resolver
}

func (r *employeeResolver) Id() *string {
	return &r.Employee.Id
}

func (r *employeeResolver) DisplayName() *string {
	return &r.Employee.DisplayName
}

func (r *employeeResolver) Email() *string {
	return &r.Employee.Email
}

func (r *employeeResolver) HireDate() *string {
	return &r.Employee.HireDate
}

func (r *employeeResolver) Department() *string {
	return &r.Employee.Department
}

func (r *employeeResolver) JobTitle() *string {
	return &r.Employee.JobTitle
}

func (r *employeeResolver) Manager() *string {
	return &r.Employee.Manager
}

func (r *employeeResolver) Practice() *string {
	return &r.Employee.Practice
}

func (r *employeeResolver) EmployeeNetworkId() *string {
	return &r.Employee.EmployeeNetworkId
}

func (r *Resolver) GetAllEmployees(ctx context.Context) (*[]*employeeResolver, error) {

	authorization := ctx.Value("X-Authorization")

	t := authorization.(string)

	array := datasources.GetAllEmployees(t)

	employeeResolvers := make([]*employeeResolver, 0)

	for _, emp := range array {
		e := models.NewEmployee(&emp)
		employeeResolvers = append(employeeResolvers, &employeeResolver{
			Employee:     e,
			rootResolver: r,
		})
	}

	return &employeeResolvers, nil

}

func (r *Resolver) GetEmployee(ctx context.Context, args struct{ ID string }) (*employeeResolver, error) {

	authorization := ctx.Value("X-Authorization")

	t := authorization.(string)

	emp := datasources.GetEmployee(t, args.ID)

	e := models.NewEmployee(&emp)

	return &employeeResolver{
		Employee:     e,
		rootResolver: r,
	}, nil

}
