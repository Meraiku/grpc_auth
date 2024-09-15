package converter

import (
	"github.com/Meraiku/grpc_auth/internal/model"
	storageModel "github.com/Meraiku/grpc_auth/internal/storage/postgres/model"
)

func ToUserFromStorage(u *storageModel.User) *model.User {
	return &model.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromUserToStorage(u *model.User) *storageModel.User {
	return &storageModel.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  []byte(u.Password),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
