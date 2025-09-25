package repository

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TbSessionName = "user_sessions"

type sessionDB struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) sessionDB {
	return sessionDB{db: db}
}
func (repo sessionDB) Create(ctx context.Context, data *domain.Session) error {
	dto := SessionDTO{
		Id:           data.Id(),
		UserId:       data.UserId(),
		RefreshToken: data.RefreshToken(),
		AccessExpAt:  data.AccessExpAt(),
		RefreshExpAt: data.RefreshExpAt(),
	}

	return repo.db.Table(TbSessionName).Create(&dto).Error
}
func (repo sessionDB) Find(ctx context.Context, id string) (*domain.Session, error) {
	var dto SessionDTO

	if err := repo.db.Table(TbSessionName).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}
func (repo sessionDB) FindByRefreshToken(ctx context.Context, rt string) (*domain.Session, error) {
	var dto SessionDTO
	if err := repo.db.Table(TbSessionName).Where("refresh_token = ?", rt).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
	}
	return dto.ToEntity()
}

func (repo sessionDB) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.Table(TbSessionName).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
func (repo sessionDB) CountSessionByUserId(ctx context.Context, userId uuid.UUID) (int64, error) {
	var count int64
	if err := repo.db.Table(TbSessionName).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
