package tokens

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {

	id := uuid.NewString()
	tokensId := uuid.NewString()
	email := gofakeit.Email()

	secret := "secret"

	token, err := GenerateJWT(
		id,
		tokensId,
		email,
		time.Minute,
		secret,
	)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := ParseJWT(token, []byte(secret))

	require.NoError(t, err)

	require.Equal(t, claims.UID, tokensId)
}

func TestJWTWithoutSecret(t *testing.T) {

	id := uuid.NewString()
	tokensId := uuid.NewString()
	email := gofakeit.Email()

	_, err := GenerateJWT(
		id,
		tokensId,
		email,
		time.Minute,
		"",
	)

	require.EqualError(t, err, ErrNoSecret.Error())

}
