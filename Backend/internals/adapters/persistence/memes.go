package db

import (
	"context"
	"database/sql"
	"memes-hustle/internals/core/memes"
)

type MemesRepoImpl struct {
	repo *sql.DB
}

func NewMemesRepo(db *sql.DB) MemesRepoImpl {
	return MemesRepoImpl{
		repo: db,
	}
}

func (r *MemesRepoImpl) InsertMemes(ctx context.Context, meme *memes.Memes) error {
	query := `INSERT INTO memes (title,image_url) VALUES($1,$2);`
	_, err := r.repo.ExecContext(ctx, query, meme.Title, meme.ImageURL)
	if err != nil {
		return err
	}
	return nil
}

func (r *MemesRepoImpl) GetAllMemes(ctx context.Context) ([]*memes.Memes, error) {
	query := `SELECT meme_id,title,image_url FROM memes;`
	allMemes := []*memes.Memes{}

	rows, err := r.repo.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRowFound
		}
		return nil, err
	}

	for rows.Next() {
		data := &memes.Memes{}
		err := rows.Scan(&data.ID, &data.Title, &data.ImageURL)
		if err != nil {
			return nil, err
		}
		allMemes = append(allMemes, data)
	}
	if len(allMemes) == 0 {
		return nil, ErrNoRowFound
	}
	return allMemes, nil
}
