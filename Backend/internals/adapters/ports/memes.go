package ports

import (
	"context"
	"memes-hustle/internals/core/memes"
)

type MemesRepository interface {
	InsertMemes(ctx context.Context, meme *memes.Memes) error
	GetAllMemes(ctx context.Context) ([]*memes.Memes, error)
}
