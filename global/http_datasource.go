package global

type HttpDataSource interface{
    Get(endpoint string) ([]byte, error)
}