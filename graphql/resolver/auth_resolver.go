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

type AuthResponse struct {
	AccessToken  string   `json:"accesstoken"`
	DisplayName  string   `json:"displayName"`
	RefreshToken string   `json:"refreshtoken"`
	Username     string   `json:"username"`
	Groups       []string `json:"groups"`
}

type authResolver struct {
    Authorization *models.AuthResponse
}

func (r *authResolver) AccessToken() *string {
	return &r.Authorization.AccessToken
}

func (r *authResolver) DisplayName() *string {
	return &r.Authorization.DisplayName
}

func (r *authResolver) RefreshToken() *string {
	return &r.Authorization.RefreshToken
}

func (r *authResolver) Username() *string {
	return &r.Authorization.Username
}

func (r *authResolver) Groups() *[]string {
	return &r.Authorization.Groups
}


func (r *Resolver) Authenticate(ctx context.Context, args struct{ Input *models.User }) (*authResolver) {
fmt.Println(args)
sampleUser:=args.Input
    xauth:=datasources.Authenticate(sampleUser)
fmt.Println(&authResolver{Authorization: &xauth})
    return &authResolver{Authorization: &xauth}

}

// func (r *Resolver) Authenticate(ctx context.Context, args *models.User) (*authResolver) {
// fmt.Println(args)
//     xauth:=datasources.Authenticate(args)
// fmt.Println(&authResolver{Authorization: &xauth})
//     return &authResolver{Authorization: &xauth}
//
// }