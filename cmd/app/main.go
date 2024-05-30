package main

import (
	"edwin-perqara-interview-test/internal/api"
	"edwin-perqara-interview-test/conf"
	"edwin-perqara-interview-test/generated"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	var server generated.ServerInterface = newServer()
	generated.RegisterHandlers(r, server)

	r.Run(fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port))
}

func newServer() *api.Server {
	opts := api.NewServerOptions{
		
	}

	return api.NewServer(opts)
}