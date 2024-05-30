package services

import "github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"

type Service struct {
	repository repository.RepositoryInterface
}

func NewService(
	repository repository.RepositoryInterface,
) *Service {
	return &Service{
		repository: repository,
	}
}
