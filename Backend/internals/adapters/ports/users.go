package ports

import (
	"context"
	"memes-hustle/internals/core/users"
)

type UserRepo interface {
	GetUser(ctx context.Context, userName string) (*users.User, error)
	InsertUser(ctx context.Context, userName string) error
}
