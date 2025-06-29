package memes

import (
	"context"
	"time"
)

type Memes struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	ImageURL  string     `json:"image_url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type MemeUseCase interface {
	InsertMeme(ctx context.Context, meme *Memes) error
	GetAllMemes(ctx context.Context) ([]*Memes, error)
}
