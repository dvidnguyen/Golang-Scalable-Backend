package repository

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"

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
func (repo sessionDb) Find(ctx context.Context, id string) (*domain.Session, error) {
	var dto SessionDTO

	if err := repo.db.Table(TbSessionName).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}
