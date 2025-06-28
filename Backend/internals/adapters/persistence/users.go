package db

import (
	"context"
	"database/sql"
	"memes-hustle/internals/core/users"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetUser(ctx context.Context, userName string) (*users.User, error) {
	query := `SELECT * FROM users where username = $1;`

	var user *users.User

	err := repo.db.QueryRowContext(ctx, query, userName).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRowFound
		}
		return nil, err
	}
	return user, nil

}
func (repo *UserRepository) InsertUser(ctx context.Context, userName string) error {
	query := `INSERT INTO users (user_id,username) VALUES ($1,$2);`
	_, err := repo.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
