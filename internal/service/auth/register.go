package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, user *model.User) (string, error) {

	defer s.log.Sync()

	const op = "Auth.Register"
	var err error

	log := s.log.With(
		zap.String("op", op),
		zap.String("email", user.Email),
	)

	log.Info("registering user")

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	user.Password, err = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash",
			zap.String("error", err.Error()),
		)

		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.storage.SaveUser(ctx, user)
	if err != nil {

		log.Error("failed to save user",
			zap.String("error", err.Error()),
		)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
