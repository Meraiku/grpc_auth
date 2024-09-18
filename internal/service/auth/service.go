package auth

import (
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/Meraiku/grpc_auth/internal/storage"
)

type service struct {
	storage storage.Storage
	log     *slog.Logger
	cfg     *config.Config
}

func NewService(
	storage storage.Storage,
	log *slog.Logger,
	cfg *config.Config,
) *service {
	return &service{
		storage: storage,
		log:     log,
		cfg:     cfg,
	}
}
