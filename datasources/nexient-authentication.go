package datasources

import (
"fmt"
"net/http"
"io/ioutil"
"encoding/json"
 "bytes"
)

type Credential struct{
Username string `json:"username"`
Password string `json:"password"`
}

// type AuthResponse struct{
// AccessToken string `json:"accesstoken"`
// DisplayName: string `json:"displayname"`
// RefreshToken: string `json:"refreshtoken"`
// Username: string `json:"username"`
// Groups: string `json:"groups"`
// }

var user Credential

func Authenticate() ([]byte){

user = Credential{Username:"cciric",Password:"Dustbumper6025!"}
json, _ := json.Marshal(user)

resp, err := http.Post("https://portal.nexient.com/gateway/api/authentication/authenticate", "application/json", bytes.NewBuffer(json))
if err != nil{
fmt.Println("error making post request, nex-auth.go")
}
defer resp.Body.Close()

 bodyBytes, _ := ioutil.ReadAll(resp.Body)

return bodyBytes
}




