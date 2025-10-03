package common

import "gorm.io/gorm"

const (
	KeyRequester = "requester"
	KeyGorm      = "gorm"
	KeyJWT       = "jwt"
)

type DBContext interface {
	GetDB() *gorm.DB
}
