package postgres

import (
	"context"
	"database/sql"
	"os"
	"strings"
	"sync"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/Meraiku/grpc_auth/internal/storage"
	"github.com/Meraiku/grpc_auth/internal/storage/postgres/converter"
	storageModel "github.com/Meraiku/grpc_auth/internal/storage/postgres/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var _ storage.UserStorage = (*postgres)(nil)
var _ storage.AppStorage = (*postgres)(nil)
var _ storage.Storage = (*postgres)(nil)

type postgres struct {
	db *bun.DB
	mu *sync.RWMutex
}

func New() (*postgres, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DB_URL"))))

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return &postgres{
		db: db,
		mu: &sync.RWMutex{}}, nil
}

func (s *postgres) SaveUser(ctx context.Context, user *model.User) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.NewInsert().Model(converter.FromUserToStorage(user)).Exec(ctx)
	if err != nil {
		tx.Rollback()

		if strings.Contains(err.Error(), "duplicate key") {
			return "", ErrEmailExists
		}

		return "", err
	}

	tx.Commit()

	return user.ID, nil
}

func (s *postgres) DeleteUser(ctx context.Context, email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.NewDelete().Table("users").Where("email = ?", email).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgres) GetUser(ctx context.Context, email string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u := &storageModel.User{}

	err := s.db.NewSelect().Model(u).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromStorage(u), nil
}

func (s *postgres) App(ctx context.Context, id int) (*model.App, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	a := &storageModel.App{}

	err := s.db.NewSelect().Table("app").Where("id = ?", id).Scan(ctx, a)
	if err != nil {
		return nil, err
	}
	return converter.ToAppFromStorage(a), nil
}
