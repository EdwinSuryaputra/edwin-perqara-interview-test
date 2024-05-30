package conf

import (
	"github.com/kelseyhightower/envconfig"

	"fmt"
)

var (
	App appConfig
	DB  dBConfig
)

func Init() {
	var cfg config
	if err := loadEnv(&cfg); err != nil {
		fmt.Println(err)
	}
}

func loadEnv(cfg *config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	syncConfig(cfg)

	return nil
}

func syncConfig(cfg *config) {
	App = cfg.App
	DB = cfg.DB
}
