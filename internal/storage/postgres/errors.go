package postgres

import "errors"

var (
	ErrEmailExists = errors.New("email already exists")
)
