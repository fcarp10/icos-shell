package openapi

import "github.com/spf13/viper"

var Client *APIClient

func Init() {
	cfg := NewConfiguration()
	cfg.Host = viper.GetString("server")
	Client = NewAPIClient(cfg)
}
