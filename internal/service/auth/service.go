package auth

import (
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/storage"
)

type service struct {
	userStorage storage.UserStorage
	appStorage  storage.AppStorage
	log         *slog.Logger
}

func NewService(
	userStorage storage.UserStorage,
	appStorage storage.AppStorage,
	log *slog.Logger,
) *service {
	return &service{
		userStorage: userStorage,
		appStorage:  appStorage,
		log:         log,
	}
}
