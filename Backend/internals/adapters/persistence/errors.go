package db

import "errors"

var (
	ErrNoRowFound error = errors.New("No Data Found")
)