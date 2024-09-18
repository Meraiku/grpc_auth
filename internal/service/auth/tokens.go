package auth

import (
	"time"

	"github.com/Meraiku/grpc_auth/internal/lib/tokens"
	"github.com/Meraiku/grpc_auth/internal/model"
)

func GenerateTokenPair(user *model.User, app *model.App, ttlAccess, ttlRefresh time.Duration) (*model.Tokens, error) {

	access, err := tokens.GenerateJWT(
		user,
		app,
		ttlAccess,
	)
	if err != nil {
		return nil, err
	}

	refresh, err := tokens.GenerateJWT(
		user,
		app,
		ttlRefresh,
	)
	if err != nil {
		return nil, err
	}

	return &model.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}
