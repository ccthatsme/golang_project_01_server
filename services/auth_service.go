package services

import (
	//"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/global_methods"
	"github.com/golang_project_01_server/graphql/models"
	"encoding/json"
	"fmt"
)

type AuthService interface {
	Auth(user models.User) *models.AuthResponse
}

type AuthDataSource struct {
	CompanyHttp global_methods.HttpDataSource
}

func NewAuthService(companyHttp global_methods.HttpDataSource) AuthService {
	return &AuthDataSource{
		CompanyHttp: companyHttp,
	}
}

 var responsiveAuth models.AuthResponse

func (ds *AuthDataSource) Auth(user models.User) *models.AuthResponse {

	result, err := ds.CompanyHttp.Post("authentication/authenticate", user)

	json.Unmarshal(result, &responsiveAuth)
	if err != nil {
		fmt.Println("line 58")
	}

	return &responsiveAuth

}
