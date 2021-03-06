package datasources

import (
	"bytes"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/golang_project_01_server/graphql/models"
	"net/http"

)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string   `json:"accesstoken"`
	DisplayName  string   `json:"displayName"`
	RefreshToken string   `json:"refreshtoken"`
	Username     string   `json:"username"`
	Groups       []string `json:"groups"`
}

type Employee struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type Project struct {
	Id     string `json:"id"`
	Active bool   `json:"active"`
	Name   string `json:"name"`
}

var allEmp []models.Employee
var allPro []models.Project

var m interface{}
var emp models.Employee
var pro models.Project
// var respAuth models.AuthResponse

// func Authenticate(user *models.User) []byte {
//
// 	userJson, _ := json.Marshal(*user)
//
// 	client := http.Client{}
//
// 	request, err := http.NewRequest("POST", "https://portal.nexient.com/gateway/api/authentication/authenticate", bytes.NewBuffer(userJson))
// 	request.Header.Set("Content-type", "application/json")
// 	request.Header.Set("Version", "2")
// 	if err != nil {
// 		fmt.Println("line 40 nex-auth")
// 	}
//
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println("error making post request, nex-auth.go")
// 	}
// 	access := resp.Header["Accesstoken"]
// 	refresh := resp.Header["Refreshtoken"]
//
// 	respAuth.AccessToken = strings.Join(access, "")
// 	respAuth.RefreshToken = strings.Join(refresh, "")
//
// 	defer resp.Body.Close()
//
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("line 57")
// 	}
//
// 	json.Unmarshal(bodyBytes, &respAuth)
// 	if err != nil {
// 		fmt.Println("line 58")
// 	}
//
// 	return bodyBytes
// }

func Authenticate(user *models.User) models.AuthResponse {

	userJson, _ := json.Marshal(*user)

	client := http.Client{}

	request, err := http.NewRequest("POST", "https://portal.nexient.com/gateway/api/authentication/authenticate", bytes.NewBuffer(userJson))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Version", "2")
	if err != nil {
		fmt.Println("line 40 nex-auth")
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error making post request, nex-auth.go")
	}
	access := resp.Header["Accesstoken"]
	refresh := resp.Header["Refreshtoken"]

	respAuth.AccessToken = strings.Join(access, "")
	respAuth.RefreshToken = strings.Join(refresh, "")

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("line 57")
	}

	json.Unmarshal(bodyBytes, &respAuth)
	if err != nil {
		fmt.Println("line 58")
	}

	return respAuth
}

func GetAllEmployees(authKey string) []models.Employee {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://portal.nexient.com/gateway/api/employees", nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Version", "2")
	request.Header.Set("X-Authorization", authKey)
	if err != nil {
		fmt.Println("line 40 nex-auth")
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error making post request, nex-auth.go")
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("line 57")
	}

	json.Unmarshal(bodyBytes, &allEmp)
	if err != nil {
		fmt.Println("line 58")
	}
	return allEmp

}

func GetEmployee(authKey string, employeeId string) models.Employee {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://portal.nexient.com/gateway/api/employees/"+employeeId, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Version", "2")
	request.Header.Set("X-Authorization", authKey)
	if err != nil {
		fmt.Println("line 40 nex-auth")
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error making post request, nex-auth.go")
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("line 57")
	}
	fmt.Println(string(bodyBytes))
	json.Unmarshal(bodyBytes, &emp)
	if err != nil {
		fmt.Println("line 58")
	}
	return emp

}

func GetAllProjects(authKey string) []models.Project {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://portal.nexient.com/gateway/api/projects", nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Version", "1")
	request.Header.Set("X-Authorization", authKey)
	if err != nil {
		fmt.Println("line 40 nex-auth")
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error making post request, nex-auth.go")
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("line 57")
	}

	json.Unmarshal(bodyBytes, &allPro)
	if err != nil {
		fmt.Println("line 58")
	}
	return allPro

}

func GetProject(authKey string, projectId string) models.Project {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://portal.nexient.com/gateway/api/projects/"+projectId, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Version", "1")
	request.Header.Set("X-Authorization", authKey)
	if err != nil {
		fmt.Println("line 40 nex-auth")
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("error making post request, nex-auth.go")
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("line 57")
	}

	json.Unmarshal(bodyBytes, &pro)
	if err != nil {
		fmt.Println("line 58")
	}
	return pro

}
