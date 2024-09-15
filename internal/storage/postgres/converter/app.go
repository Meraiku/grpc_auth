package converter

import (
	"github.com/Meraiku/grpc_auth/internal/model"
	storageModel "github.com/Meraiku/grpc_auth/internal/storage/postgres/model"
)

func ToAppFromStorage(a *storageModel.App) *model.App {
	return &model.App{
		ID:     a.ID,
		Name:   a.Name,
		Secret: a.Secret,
	}
}
