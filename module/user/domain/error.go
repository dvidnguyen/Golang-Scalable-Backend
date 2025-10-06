package domain

import "errors"

var (
	ErrEmailHasExisted    = errors.New("email already exists")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrUserBanned         = errors.New("user is banned")
	ErrTooManyLogin       = errors.New("too many user login")
	ErrCannotChangeAvatar = errors.New("cannot change avatar of user")
)
