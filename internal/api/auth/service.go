package auth

import (
	"github.com/Meraiku/grpc_auth/internal/service"
	ssov1 "github.com/Meraiku/protos/gen/go/sso"
)

type Implemintation struct {
	ssov1.UnimplementedAuthServer
	authService service.AuthService
}

func NewImplemintation(authService service.AuthService) *Implemintation {
	return &Implemintation{
		authService: authService,
	}
}
