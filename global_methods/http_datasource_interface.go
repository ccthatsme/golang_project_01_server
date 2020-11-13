package global_methods

type HttpDataSource interface {
	Get(endpoint string, authKey string) ([]byte, error)
	Post(endpoint string, data interface{}) ([]byte, error)
}