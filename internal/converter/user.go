package converter

import (
	"github.com/Meraiku/grpc_auth/internal/model"
	ssov1 "github.com/Meraiku/protos/gen/go/sso"
)

func ToUserFromSSO(user *ssov1.RegisterRequest) *model.User {
	return &model.User{
		Email:    user.Email,
		Password: []byte(user.Password),
	}
}
