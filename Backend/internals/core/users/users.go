package users

import (
	"context"
	"time"
)

type User struct {
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserUsecase interface {
	InsertUser(ctx context.Context, userName string) error
	GetUserByUserName(ctx context.Context, userName string) (*User, error)
}
