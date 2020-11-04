package resolvers

import (
	"encoding/json"
	"fmt"
	//"github.com/friendsofgo/graphiql"
	//"github.com/golang_project_01_server/datasources"
// 	"github.com/golang_project_01_server/services"
	//"github.com/gorilla/mux"
	//"github.com/graphql-go/graphql"
	//"io/ioutil"
	"net/http"
	"context"
	"github.com/golang_project_01_server/Models"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() queryResolver {
	return &queryResolver{r}
}

func (r *Resolver) GetAllEmployees(ctx context.Context()) (*[]Models.Employee, err) {

    employees, err := r.EmployeeService.GetAllEmployees()
	if err != nil {
	    fmt.Println("Error in employee resolver")
	        return &Models.Employees
	}

            return
}
