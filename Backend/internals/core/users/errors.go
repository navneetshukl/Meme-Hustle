package users

import "errors"

var (
	ErrUserAlreadyExist error = errors.New("username already exists")
)