package conf

func InitTest() {
	App.Name = "edwin-perqara-interview-test"
	App.Port = 5000

	DB.Host = "localhost"
	DB.Port = 5432
	DB.User = "postgres"
	DB.Pass = "12345678"
	DB.Name = "postgres"
}
