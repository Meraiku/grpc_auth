package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Storage struct {
	db *bun.DB
}

func New(dsn string) (*Storage, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (string, error) {
	u := &User{
		ID:       uuid.New(),
		Email:    email,
		Password: passHash,
	}

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

	return u.ID.String(), nil
}

func (s *Storage) DeleteUser(ctx context.Context, email string) error {

	_, err := s.db.NewDelete().Table("users").Where("email = ?", email).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) User(ctx context.Context, email string) (*User, error) {
	u := &User{}

	_, err := s.db.NewSelect().Model(u).Where("email = ?", email).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Storage) App(ctx context.Context, id int) (*App, error) {
	a := &App{}

	_, err := s.db.NewSelect().Model(a).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}
