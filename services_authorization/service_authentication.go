package services_authorization

import (
	"encoding/json"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
	"net/http"
)

type AuthResponse struct {
	AccessToken  string   `json:"accesstoken"`
	DisplayName  string   `json:"displayName"`
	RefreshToken string   `json:"refreshtoken"`
	Username     string   `json:"username"`
	Groups       []string `json:"groups"`
}

type XAuth struct {
	Token string `json:"token"`
}

var resp AuthResponse

var AuthToken XAuth
var AuthTokenTwo XAuth

func ServiceAuthorization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")

	decoder := json.NewDecoder(r.Body)
	// 	user := &datasources.Credential{}
	user := &models.User{}
	err := decoder.Decode(user)
	if err != nil {
		panic(err)
	}

	array := datasources.Authenticate(user)

	AuthToken.Token = array.AccessToken

	json.NewEncoder(w).Encode(array)

}
