package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Meraiku/grpc_auth/internal/domain/models"
	"github.com/Meraiku/grpc_auth/internal/lib/jwt"
	"github.com/Meraiku/grpc_auth/internal/lib/logger/sl"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid string, err error)
	User(ctx context.Context, email string) (models.User, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

type Auth struct {
	log         *slog.Logger
	userStorage UserStorage
	appProvider AppProvider
	tokenTTL    time.Duration
}

func New(log *slog.Logger, userStorage UserStorage, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		userStorage: userStorage,
		log:         log,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (string, error) {

	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userStorage.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (*jwt.Tokens, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	user, err := a.userStorage.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	tok := jwt.NewTokens()

	access, err := tok.ID(user.ID).Email(user.Email).ExpiredAt(15 * time.Minute).AppID(app.ID).Generate([]byte(app.Secret))
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	refresh, err := tok.ID(user.ID).ExpiredAt(24 * time.Hour).AppID(app.ID).Generate([]byte(app.Secret))
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &jwt.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}
