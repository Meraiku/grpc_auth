package tokens

import (
	"testing"
	"time"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {

	u := &model.User{
		ID:    uuid.NewString(),
		Email: gofakeit.Email(),
	}

	a := &model.App{
		ID:     1,
		Secret: "secret",
	}

	token, err := GenerateJWT(
		u,
		a,
		time.Minute,
	)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := ParseJWT(token, []byte(a.Secret))

	require.NoError(t, err)

	require.Equal(t, claims.ID, u.ID)
}

func TestJWTWithoutSecret(t *testing.T) {

	u := &model.User{
		ID:    uuid.NewString(),
		Email: gofakeit.Email(),
	}

	a := &model.App{
		ID: 1,
	}

	_, err := GenerateJWT(
		u,
		a,
		time.Minute,
	)

	require.EqualError(t, err, ErrNoSecret.Error())

}

func TestJWTParsing(t *testing.T) {

	u := &model.User{
		ID:    uuid.NewString(),
		Email: gofakeit.Email(),
	}

	a := &model.App{
		ID:     1,
		Secret: "secret",
	}

	token, err := GenerateJWT(
		u,
		a,
		time.Minute,
	)

	require.NoError(t, err)

	_, err = ParseJWT(token, []byte(""))

	require.Error(t, err)
}
