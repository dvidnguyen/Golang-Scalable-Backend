package repository

import (
	"Ls04_GORM/module/user/domain"
	"context"

	"gorm.io/gorm"
)

const TbSessionName = "user_sessions"

type sessionDb struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) sessionDb {
	return sessionDb{db: db}
}
func (repo sessionDb) Create(ctx context.Context, data *domain.Session) error {
	dto := SessionDTO{
		Id:           data.Id(),
		UserId:       data.UserId(),
		RefreshToken: data.RefreshToken(),
		AccessExpAt:  data.AccessExpAt(),
		RefreshExpAt: data.RefreshExpAt(),
	}

	return repo.db.Table(TbSessionName).Create(&dto).Error
}
