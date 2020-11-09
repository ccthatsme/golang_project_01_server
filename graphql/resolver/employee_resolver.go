package resolver

import (
	"fmt"
	"context"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)



type employeeResolver struct {
	Employee *models.Employee
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

func (r *Resolver) GetAllEmployees(ctx context.Context) (*[]*employeeResolver, error) {

	authorization := ctx.Value("X-Authorization")
	t := authorization.(string)
	fmt.Println( "-----------------------------------------")
	fmt.Println(t)
	array := datasources.GetAllEmployees(t)

	employeeResolvers := make([]*employeeResolver, 0)

	for _, emp := range array {
		employeeResolvers = append(employeeResolvers, &employeeResolver{
			Employee: &emp,
// 			rootResolver: r,
		})
	}

	return &employeeResolvers, nil

}
