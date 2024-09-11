package main

import (
	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	_ = log

	db, err := connectDB()
	if err != nil {
		log.Error(err.Error())
	}

	_ = db
}
