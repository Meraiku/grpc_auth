package service

import (
	"context"

	"github.com/Meraiku/grpc_auth/internal/model"
)

type AuthService interface {
	Login(ctx context.Context, user *model.User, appID int) (*model.Tokens, error)
	Register(ctx context.Context, user *model.User) (string, error)
}
