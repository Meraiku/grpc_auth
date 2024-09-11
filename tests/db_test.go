package tests

import (
	"context"
	"os"
	"testing"

	"github.com/Meraiku/grpc_auth/internal/storage/postgres"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	godotenv.Load(".env")

	db, err := postgres.New(os.Getenv("DB_URL"))

	if err != nil {
		t.Skip(err)
	}

	id, err := db.SaveUser(context.Background(), "test@gmail.com", []byte("pass"))

	if !assert.Nil(t, err, "want user creation, but got err: %s", err) {
		t.Fail()
	}

	err = uuid.Validate(id)
	if !assert.Nil(t, err, "want valid uuid, but got err: %s", err) {
		t.Fail()
	}

}
