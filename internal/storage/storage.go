package storage

import (
	"context"
	"errors"

	"github.com/Meraiku/grpc_auth/internal/model"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)

type UserStorage interface {
	SaveUser(ctx context.Context, user *model.User) (string, error)
	GetUser(ctx context.Context, email string) (*model.User, error)
}

type AppStorage interface {
	App(ctx context.Context, id int) (*model.App, error)
}

type Storage interface {
	UserStorage
	AppStorage
}
