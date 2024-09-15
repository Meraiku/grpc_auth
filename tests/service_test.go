package tests

import (
	"context"
	"testing"

	"github.com/Meraiku/grpc_auth/internal/config"
	ssov1 "github.com/Meraiku/protos/gen/go/sso"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test_secret"
)

type Suit struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suit) {
	t.Helper()
	t.Parallel()

	config.Load(".env")

	cfg, err := config.NewGRPCConfig()
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	cc, err := grpc.NewClient(
		cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %s", err)
	}

	return ctx, &Suit{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func TestGRPCService(t *testing.T) {
	ctx, st := New(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})

	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLog, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, respLog.Access)
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}
