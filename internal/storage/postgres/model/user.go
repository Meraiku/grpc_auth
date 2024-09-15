package model

import "time"

type User struct {
	ID        string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
