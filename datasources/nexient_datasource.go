package datasources

import (
	"github.com/golang_project_01_server/global_methods"
	"github.com/golang_project_01_server/graphql/models"
	"encoding/json"
	"net/http"
	"fmt"
		"bytes"
    	"strings"
    	"io/ioutil"
)

type NexientDataSource struct {
	BaseEndPoint string
}

func NewNexientDataSource(baseEndPoint string) (global_methods.HttpDataSource, error) {

	return &NexientDataSource{
		BaseEndPoint: baseEndPoint,
	}, nil
}

var respAuth models.AuthResponse

func (ds *NexientDataSource) Get(endPoint string) ([]byte, error) {

	return make([]byte, 2), nil
}

func (ds *NexientDataSource) Post(endpoint string, data interface{}) ([]byte, error) {

	userJson, _ := json.Marshal(&data)

	client := http.Client{}

	request, err := http.NewRequest("POST", "https://portal.nexient.com/gateway/api/"+endpoint, bytes.NewBuffer(userJson))
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
fmt.Println(string(bodyBytes[0]), "datasource")
	return bodyBytes, nil

}
