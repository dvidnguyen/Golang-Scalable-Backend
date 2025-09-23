package domain

import "errors"

var (
	ErrEmailHasExisted = errors.New("email already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserBanned      = errors.New("user is banned")
)
