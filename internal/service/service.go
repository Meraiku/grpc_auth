package service

import (
	"context"

	"github.com/Meraiku/grpc_auth/internal/lib/jwt"
	"github.com/Meraiku/grpc_auth/internal/model"
)

type AuthService interface {
	Login(ctx context.Context, user *model.User, appID int) (*jwt.Tokens, error)
	Register(ctx context.Context, user *model.User) (string, error)
}
