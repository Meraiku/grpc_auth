package converter

import (
	"github.com/Meraiku/grpc_auth/internal/model"
	ssov1 "github.com/Meraiku/protos/gen/go/sso"
)

func ToUserFromSSORegister(user *ssov1.RegisterRequest) *model.User {
	return &model.User{
		Email:    user.Email,
		Password: []byte(user.Password),
	}
}

func ToUserFromSSOLogin(user *ssov1.LoginRequest) *model.User {
	return &model.User{
		Email:    user.Email,
		Password: []byte(user.Password),
	}
}

func ToAppFromSSOLogin(app *ssov1.LoginRequest) *model.App {
	return &model.App{
		ID: int(app.AppId),
	}
}
