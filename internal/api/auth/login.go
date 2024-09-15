package auth

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/converter"
	"github.com/Meraiku/grpc_auth/internal/service/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ssov1 "github.com/Meraiku/protos/gen/go/sso"
)

func (i *Implemintation) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if in.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	tokens, err := i.authService.Login(ctx, converter.ToUserFromSSOLogin(in), converter.ToAppFromSSOLogin(in).ID)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		slog.Error("error login user", "error", err)

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Access: tokens.AccessToken, Refresh: tokens.RefreshToken}, nil
}
