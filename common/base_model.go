package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"column:id;"`
	Status    string    `gorm:"column:status;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func GenNewModel() BaseModel {
	now := time.Now().UTC()
	newId, _ := uuid.NewV7()

	return BaseModel{
		Id:        newId,
		Status:    "activated",
		CreatedAt: now,
		UpdatedAt: now,
	}
}
