package app

import (
	"log"
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/api/auth"
	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/service"
	authService "github.com/Meraiku/grpc_auth/internal/service/auth"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"github.com/Meraiku/grpc_auth/internal/storage/postgres"
)

type serviceProvider struct {
	log         *slog.Logger
	config      *config.Config
	storage     storage.Storage
	authService service.AuthService
	authImpl    *auth.Implemintation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err)
		}
		s.config = cfg
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

func (s *serviceProvider) Logger() *slog.Logger {
	if s.log == nil {
		s.log = sl.SetupLogger(s.Config().Env)
	}

	return s.log
}

func (s *serviceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.Storage(), s.Logger())
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl() *auth.Implemintation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplemintation(s.AuthService())
	}

	return s.authImpl
}
