package auth

import (
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/storage"
)

type service struct {
	storage storage.Storage
	log     *slog.Logger
}

func NewService(
	storage storage.Storage,
	log *slog.Logger,
) *service {
	return &service{
		storage: storage,
		log:     log,
	}
}
