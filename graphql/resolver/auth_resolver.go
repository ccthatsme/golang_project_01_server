package resolver

import (
	"fmt"
	"context"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

// type User struct{
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

type authResolver struct {
    Authorization *models.AuthResponse
}

func (r *authResolver) AccessToken() *string {
	return &r.Authorization.AccessToken
}

func (r *Resolver) Authenticate(ctx context.Context, args *models.User) (*authResolver) {
fmt.Println(args)
    xauth:=datasources.Authenticate(args)
fmt.Println(&authResolver{Authorization: &xauth})
    return &authResolver{Authorization: &xauth}

}