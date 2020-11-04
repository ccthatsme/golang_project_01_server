package datasources

import (
	"github.com/golang_project_01_server/Models"
	"github.com/golang_project_01_server/global"
)

type NexientDataSource struct {
	BaseEndpoint string
}

func NewCompanyDatasource(baseEndPoint string) *global.HttpDataSource {
	return &NexientDataSource{
		BaseEndpoint: baseEndPoint,
	}
}

func (ds *NexientDataSource) Get(endpoint string) ([]byte, error) {

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
		fmt.Println("")
	}

	return bodyBytes, nil
}
