package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, user *model.User) (string, error) {

	const op = "Auth.RegisterNewUser"
	var err error

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", user.Email),
	)

	log.Info("registering user")

	user.Password, err = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.userStorage.SaveUser(ctx, user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
