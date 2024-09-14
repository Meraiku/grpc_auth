package postgres

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password []byte    `json:"password"`
}

type App struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
