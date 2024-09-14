package postgres

import (
	"context"
	"database/sql"
	"sync"

	"github.com/Meraiku/grpc_auth/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type postgres struct {
	db *bun.DB
	mu *sync.RWMutex
}

func New(dsn string) (*postgres, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &postgres{
		db: db,
		mu: &sync.RWMutex{}}, nil
}

func (s *postgres) SaveUser(ctx context.Context, u *model.User) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.NewInsert().Model(u).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()

	return u.ID, nil
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
	u := &model.User{}

	_, err := s.db.NewSelect().Model(u).Where("email = ?", email).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *postgres) App(ctx context.Context, id int) (*model.App, error) {
	a := &model.App{}

	_, err := s.db.NewSelect().Model(a).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}
