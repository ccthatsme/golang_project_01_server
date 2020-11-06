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
	array := datasources.GetAllEmployees("eyJhbGciOiJIUzUxMiIsInppcCI6IkRFRiJ9.eNosi0EKgzAQRe8yayOaJpPgvotuW3oAM44wQqOYlAqld28CLt9_739hyQIDkCNjyQY1OzMpo9EpfyFWwXlvtLYzBgsNpHeoMckuVJCPDYYeO4Md9tY3IClVv77ayIdwzO3yyfVH68bFPB_Xe8HMcYz5NpXl7OD3BwAA__8.-VywZJf_czidl8rHTRoWLUepR6A_F5AKq8NDmo1UxtNCGgX73nNpX1aBvH8nosWIgMKF2hSD3qITPjl8B_25Bg")

	employeeResolvers := make([]*employeeResolver, 0)

	for _, emp := range array {
		employeeResolvers = append(employeeResolvers, &employeeResolver{
			Employee: &emp,
// 			rootResolver: r,
		})
	}

	return &employeeResolvers, nil

}
