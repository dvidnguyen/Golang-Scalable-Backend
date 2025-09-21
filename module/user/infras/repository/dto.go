package repository

import (
	"Ls04_GORM/module/user/domain"

	"github.com/google/uuid"
)

type UserDTO struct {
	Id        uuid.UUID
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Salt      string `gorm:"column:salt" json:"salt"`
	Role      string `gorm:"column:role" json:"role"`
}

func (dto UserDTO) ToEntity() (*domain.User, error) {
	return domain.NewUser(dto.Id, dto.FirstName, dto.LastName, dto.Email, dto.Password, dto.Salt, domain.GetRole(dto.Role))
}
