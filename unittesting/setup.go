package unittesting

import "github.com/naphat-sirisubkulchai/shop/config"

func NewTestConfig() *config.Config {
	cfg := config.LoadConfig("../env/test/.env")
	return &cfg
}