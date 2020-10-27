package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang_project_01_server/datasources"
	"net/http"
	//"github.com/golang_project_01_server/routes/middleware"
	"github.com/gorilla/mux"
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

	AuthToken.Token = array.AccessToken

	json.NewEncoder(w).Encode(array)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")

	array := datasources.GetAllEmployees(AuthToken.Token)

	json.NewEncoder(w).Encode(array)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Version", "2")
	params := mux.Vars(r)
	fmt.Println(params["employeeNetworkid"])
	array := datasources.GetEmployee(AuthToken.Token, params["employeeNetworkid"])

	json.NewEncoder(w).Encode(array)
}

//middleware for checking token
func CheckTokenExists(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "X-Authorization", AuthToken)
		fmt.Println(ctx, "this is the context")

		k := ctx.Value("X-Authorization")
		fmt.Println(k)
		if k != nil {
			next(w, r.WithContext(ctx))
		}
	})
}
