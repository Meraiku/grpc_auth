package app

import (
	"log"

	"github.com/Meraiku/grpc_auth/internal/api/auth"
	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/Meraiku/grpc_auth/internal/service"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"github.com/Meraiku/grpc_auth/internal/storage/postgres"
)

type serviceProvider struct {
	config      *config.Config
	storage     storage.Storage
	authService *service.AuthService
	authImpl    *auth.Implemintation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		s.config = config.MustLoad()
	}

	return s.config
}

func (s *serviceProvider) Storage() storage.Storage {
	if s.storage == nil {
		db, err := postgres.New()
		if err != nil {
			log.Fatalf("error connecting DB: %s", err)
		}
		s.storage = db
	}

	return s.storage
}

func (s *serviceProvider) AuthService() *service.AuthService {
	if s.authService == nil {
		s.authService = 
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl() *auth.Implemintation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplemintation(*s.authService)
	}

	return s.authImpl
}
