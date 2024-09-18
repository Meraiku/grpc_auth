package service

import (
	"context"

	"github.com/Meraiku/grpc_auth/internal/lib/tokens"
	"github.com/Meraiku/grpc_auth/internal/model"
)

type AuthService interface {
	Login(ctx context.Context, user *model.User, appID int) (*tokens.Tokens, error)
	Register(ctx context.Context, user *model.User) (string, error)
}
