package services

import (
"fmt"
	"encoding/json"
	"github.com/golang_project_01_server/datasources"
	"net/http"
	//"github.com/golang_project_01_server/routes/middleware"
)

type AuthResponse struct {
	AccessToken  string   `json:"accesstoken"`
	DisplayName  string   `json:"displayName"`
	RefreshToken string   `json:"refreshtoken"`
	Username     string   `json:"username"`
	Groups       []string `json:"groups"`
}

var resp AuthResponse

var token string

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")

	decoder := json.NewDecoder(r.Body)
	user := &datasources.Credential{}
	err := decoder.Decode(user)
	if err != nil {
		panic(err)
	}

	array := datasources.Authenticate(user)

token = array.AccessToken
fmt.Println("___________________test")
fmt.Println(array.AccessToken)
fmt.Printf("%T",array.AccessToken)

	json.NewEncoder(w).Encode(array)

}

//middleware for checking token
func CheckTokenExists(next http.HandlerFunc) http.HandlerFunc{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
ctx := r.Context()
fmt.Println(ctx, "this is the context")
})
}