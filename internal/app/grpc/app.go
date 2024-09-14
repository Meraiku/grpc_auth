package grpcapp

import (
	"log/slog"

	"github.com/Meraiku/grpc_auth/internal/grpc/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

type Application struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, authService AuthService, port int) *Application {

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))
	auth.Register(grpcServer, authService)

	return &Application{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}
