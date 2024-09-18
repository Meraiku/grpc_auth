package auth

import (
	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"go.uber.org/zap"
)

type service struct {
	storage storage.Storage
	log     *zap.Logger
	cfg     *config.Config
}

func NewService(
	storage storage.Storage,
	log *zap.Logger,
	cfg *config.Config,
) *service {
	return &service{
		storage: storage,
		log:     log,
		cfg:     cfg,
	}
}
