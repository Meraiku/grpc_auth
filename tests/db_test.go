package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/Meraiku/grpc_auth/internal/storage/postgres"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		t.Skip(err)
	}

	db, err := postgres.New()

	if err != nil {
		t.Skip(err)
	}

	user := &model.User{
		ID:        uuid.NewString(),
		Email:     "test@gmail.com",
		Password:  []byte("pass"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := db.SaveUser(context.Background(), user)

	if !assert.Nil(t, err, "want user creation, but got err: %s", err) {
		t.Fatal()
	}

	err = uuid.Validate(id)
	if !assert.Nil(t, err, "want valid uuid, but got err: %s", err) {
		t.Fatal()
	}

	err = db.DeleteUser(context.Background(), user.Email)
	if !assert.Nil(t, err, "want succses on deletion user, but got err: %s", err) {
		t.Fatal()
	}
}
