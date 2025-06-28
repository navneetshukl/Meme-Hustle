package users

import (
	"context"
	db "memes-hustle/internals/adapters/persistence"
	"memes-hustle/internals/adapters/ports"

	"memes-hustle/internals/core/users"
)

type UserUseCaseImpl struct {
	userRepo ports.UserRepo
}

func NewUserUseCaseImpl(repo ports.UserRepo) users.UserUsecase {
	return &UserUseCaseImpl{
		userRepo: repo,
	}
}

func (u *UserUseCaseImpl) InsertUser(ctx context.Context, userName string) error {

	user, err := u.userRepo.GetUser(ctx, userName)
	if err != nil {
		if err != db.ErrNoRowFound {
			return err
		}
	}

	if user != nil {
		return users.ErrUserAlreadyExist
	}

	err = u.userRepo.InsertUser(ctx, userName)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCaseImpl) GetUserByUserName(ctx context.Context, userName string) (*users.User, error) {
	user, err := u.userRepo.GetUser(ctx, userName)
	if err != nil {
		if err != db.ErrNoRowFound {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}
