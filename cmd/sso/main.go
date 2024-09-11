package main

import (
	"context"

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

	_, err = db.SaveUser(context.Background(), "test@gmail.com", []byte("pass"))
	if err != nil {
		log.Error(err.Error())
	}

	if err := db.DeleteUser(context.Background(), "test@gmail.com"); err != nil {
		log.Error(err.Error())
	}
}
