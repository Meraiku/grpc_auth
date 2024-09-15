package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Meraiku/grpc_auth/internal/lib/jwt"
	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, u *model.User, appID int) (*jwt.Tokens, error) {
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

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(user.Password)); err != nil {
		s.log.Info("invalid credentials", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := s.storage.App(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	tok := jwt.NewTokens()

	access, err := tok.ID(user.ID).Email(user.Email).ExpiredAt(15 * time.Minute).AppID(app.ID).Generate([]byte(app.Secret))
	if err != nil {
		s.log.Error("failed to generate token", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	refresh, err := tok.ID(user.ID).ExpiredAt(24 * time.Hour).AppID(app.ID).Generate([]byte(app.Secret))
	if err != nil {
		s.log.Error("failed to generate token", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &jwt.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}
