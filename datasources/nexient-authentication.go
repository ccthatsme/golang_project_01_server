package datasources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user Credential

func Authenticate() ([]byte, []string) {

	user = Credential{Username: "cciric", Password: "Dustbumper6025!\""}
	json, _ := json.Marshal(user)

	client := http.Client{}

	request, err := http.NewRequest("POST", "https://portal.nexient.com/gateway/api/authentication/authenticate", bytes.NewBuffer(json))
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
	fmt.Println(access)

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyBytes))

	return bodyBytes, access
}
