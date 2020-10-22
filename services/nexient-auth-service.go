package services

import (
"net/http"
"encoding/json"
 "github.com/golang_project_01_server/datasources"
)

type AuthResponse struct{
//AccessToken string `json:"accesstoken"`
DisplayName string `json:"displayName"`
//RefreshToken string `json:"refreshtoken"`
Username string `json:"username"`
Groups string `json:"groups"`
}

var resp AuthResponse

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")

array := datasources.Authenticate()

json.Unmarshal(array, &resp)

	json.NewEncoder(w).Encode(resp)

}