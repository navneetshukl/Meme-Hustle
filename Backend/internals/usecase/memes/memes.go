package memes

import (
	"context"
	db "memes-hustle/internals/adapters/persistence"
	"memes-hustle/internals/adapters/ports"
	"memes-hustle/internals/core/memes"
)

type MemesUseCaseImpl struct {
	repo ports.MemesRepository
}

func NewMemesRepo(r ports.MemesRepository) memes.MemeUseCase {
	return &MemesUseCaseImpl{
		repo: r,
	}
}

func (r *MemesUseCaseImpl) InsertMeme(ctx context.Context, meme *memes.Memes) error {

	err := r.repo.InsertMemes(ctx, meme)
	if err != nil {
		return memes.ErrInsertingMemes
	}
	return nil
}

func(r *MemesUseCaseImpl)GetAllMemes(ctx context.Context)([]*memes.Memes,error){
	allMemes,err:=r.repo.GetAllMemes(ctx)
	if err!=nil{
		if err==db.ErrNoRowFound{
			return []*memes.Memes{},nil
		}
		return nil,err
	}
	return allMemes,nil
}
