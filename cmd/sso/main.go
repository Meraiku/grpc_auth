package main

import "github.com/Meraiku/grpc_auth/internal/config"

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	_ = log
}
