package repository

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

const TbName = "users"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{db: db}
}

func (repo userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var dto UserDTO

	if err := repo.db.Table(TbName).Where("email = ?", email).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToEntity()
}
func (repo userRepository) Create(ctx context.Context, data *domain.User) error {
	dto := UserDTO{
		Id:        data.Id(),
		FirstName: data.FirstName(),
		LastName:  data.LastName(),
		Email:     data.Email(),
		Password:  data.Password(),
		Salt:      data.Salt(),
		Role:      data.Role().String(),
	}

	if err := repo.db.Table(TbName).Create(&dto).Error; err != nil {
		return err
	}
	return nil
}
