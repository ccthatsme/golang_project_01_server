package datasources

import (
	"github.com/golang_project_01_server/global_methods"
	//"github.com/golang_project_01_server/datasources"
)

type NexientDataSource struct {
	BaseEndPoint string
}

func NewNexientDataSource(baseEndPoint string) global_methods.HttpDataSource {

	return &NexientDataSource{
		BaseEndPoint: baseEndPoint,
	}
}

func (ds *NexientDataSource) Get(endPoint string) ([]byte, error) {



}