package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, user *model.User) (string, error) {

	const op = "Auth.Register"
	var err error

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", user.Email),
	)

	log.Info("registering user")

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

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
