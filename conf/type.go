package conf

type config struct {
	App appConfig
	DB  dBConfig
}

type appConfig struct {
	Name string `envconfig:"APP_NAME"`
	Host string `envconfig:"APP_HOST"`
	Port int    `envconfig:"APP_PORT"`
}

type dBConfig struct {
	Host string `envconfig:"DB_HOST"`
	Port int    `envconfig:"DB_PORT"`
	User string `envconfig:"DB_USER"`
	Pass string `envconfig:"DB_PASS"`
	Name string `envconfig:"DB_NAME"`
}
