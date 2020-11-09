package resolver

import (
	"fmt"
	"context"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

// type AuthResponse struct {
// 	AccessToken  string   `json:"accesstoken"`
// 	DisplayName  string   `json:"displayName"`
// 	RefreshToken string   `json:"refreshtoken"`
// 	Username     string   `json:"username"`
// 	Groups       []string `json:"groups"`
// }

type XAuth struct {
	Token string `json:"token"`
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

var AuthToken models.XAuth

func (r *Resolver) Authenticate(ctx context.Context, args struct{ Input *models.User }) (*authResolver) {

    sampleUser:=args.Input
    xauth:=datasources.Authenticate(sampleUser)

    AuthToken.Token = xauth.AccessToken

    ctx = context.WithValue(ctx, "X-Authorization", xauth.AccessToken)
    fmt.Println(ctx)



// fmt.Println(&authResolver{Authorization: &xauth})

    return &authResolver{Authorization: &xauth}

}

// func (r *Resolver) Authenticate(ctx context.Context, args *models.User) (*authResolver) {
// fmt.Println(args)
//     xauth:=datasources.Authenticate(args)
// fmt.Println(&authResolver{Authorization: &xauth})
//     return &authResolver{Authorization: &xauth}
//
// }