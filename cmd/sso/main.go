package main

import (
	"context"
	"log"

	"github.com/Meraiku/grpc_auth/internal/app"
)

func main() {
	a, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("failed to init app: %s", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
