package openapi

var Client *APIClient

func Init(server string) {
	cfg := NewConfiguration()
	cfg.Host = server
	Client = NewAPIClient(cfg)
}
