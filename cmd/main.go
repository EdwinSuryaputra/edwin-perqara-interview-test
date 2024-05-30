package main

import (
	"fmt"
	"os"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/conf"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/generated"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/handler"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/EdwinSuryaputra/edwin-perqara-interview-test/docs"
)

//	@title			edwin-perqara-interview-test
//	@version		0.1.0
//	@description	This is Perqara interview purpose.

//	@host	localhost:5000
func main() {
	conf.Init()

	cfg := conf.App
	r := gin.Default()

	var server generated.ServerInterface = newServer()
	generated.RegisterHandlers(r, server)

	r.Use(middleware.OapiRequestValidator(newSwagger(r)))

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

func newSwagger(r *gin.Engine) *openapi3.T {
	swagger, err := generated.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	return swagger
}
