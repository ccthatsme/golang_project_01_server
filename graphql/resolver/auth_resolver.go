package resolver

import (
	"context"
	//"fmt"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

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

func (r *Resolver) Authenticate(ctx context.Context, args struct{ Input *models.User }) *authResolver {

	sampleUser := args.Input

	xauth := datasources.Authenticate(sampleUser)

	AuthToken.Token = xauth.AccessToken

	return &authResolver{Authorization: &xauth}
}
