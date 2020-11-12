package global_methods

type HttpDataSource interface {
	Get(endpoint string) ([]byte, error)
	//PostAuth(endpoint string, data interface) ([]byte, error)
	Post(endpoint string, data interface) ([]byte, error)
}