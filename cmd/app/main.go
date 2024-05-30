package main

import (
	"fmt"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/conf"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/generated"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/handler"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := conf.App
	r := gin.Default()

	var server generated.ServerInterface = newServer()
	generated.RegisterHandlers(r, server)

	r.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
}

func newServer() *handler.Server {
	cfg := conf.DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dsn,
	})

	var service services.ServiceInterface = services.NewService(repo)

	opts := handler.NewServerOptions{
		Repository: repo,
		Service:    service,
	}

	return handler.NewServer(opts)
}
