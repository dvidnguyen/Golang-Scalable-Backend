package domain

import "errors"

var (
	ErrEmailHasExisted = errors.New("email already exists")
)
