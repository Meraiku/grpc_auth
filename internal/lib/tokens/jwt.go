package tokens

import (
	"errors"
	"time"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrNoSecret = errors.New("error: empty secret")
)

type Claims struct {
	ID    string `json:"id"`
	UID   string `json:"uid"`
	Email string `json:"email"`
	AppID int    `json:"app_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(
	user *model.User,
	app *model.App,
	ttl time.Duration,
) (string, error) {
	if app.Secret == "" {
		return "", ErrNoSecret
	}

	c := &Claims{
		ID:    user.ID,
		UID:   uuid.NewString(),
		Email: user.Email,
		AppID: app.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl).UTC()),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	token, err := jwtToken.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseJWT(tokenStr string, secret []byte) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return secret, nil })
	if token != nil {
		if token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
