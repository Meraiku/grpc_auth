package converter

import (
	"github.com/Meraiku/grpc_auth/internal/model"
	ssov1 "github.com/Meraiku/protos/gen/go/sso"
	"github.com/google/uuid"
)

func ToUserFromSSO(user *ssov1.RegisterRequest) *model.User {
	return &model.User{
		ID:       uuid.NewString(),
		Email:    user.Email,
		Password: []byte(user.Password),
	}
}
