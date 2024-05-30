package handler

import (
	"github.com/EdwinSuryaputra/edwin-sawitpro-be-interview-test/repository"
	"github.com/EdwinSuryaputra/edwin-sawitpro-be-interview-test/services"
)

type Server struct {
	Service    services.ServiceInterface
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
	Service    services.ServiceInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Service:    opts.Service,
		Repository: opts.Repository,
	}
}
