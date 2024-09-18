package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, u *model.User, appID int) (*model.Tokens, error) {
	const op = "Auth.Login"

	log := s.log.With(
		slog.String("op", op),
		slog.String("username", u.Email),
	)

	log.Info("attempting to login user")

	user, err := s.storage.GetUser(ctx, u.Email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			s.log.Warn("user not found", sl.Err(err))

			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		s.log.Error("failed to get user", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {
		s.log.Info("invalid credentials", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := s.storage.App(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	tokenPair, err := GenerateTokenPair(u, app, s.cfg.AccessTTL, s.cfg.RefreshTTL)
	if err != nil {
		s.log.Error("failed to generate tokens", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tokenPair, nil
}
