package query

import (
	"Ls04_GORM/common"
	"context"

	"github.com/google/uuid"
	sctx "github.com/viettranx/service-context"
)

type CategoryDTO struct {
	Id    uuid.UUID `gorm:"column:id" json:"id"`
	Title string    `gorm:"column:title" json:"title"`
}

func (CategoryDTO) TableName() string {
	return "categories"
}

type CategoryById struct {
	sctx sctx.ServiceContext
}

func NewCategoryById(sctx sctx.ServiceContext) *CategoryById {
	return &CategoryById{sctx: sctx}
}
func (cat *CategoryById) Execute(ctx context.Context, ids []uuid.UUID) ([]CategoryDTO, error) {
	var categories []CategoryDTO
	dbContext := cat.sctx.MustGet(common.KeyGorm).(common.DBContext)

	if err := dbContext.GetDB().Table(CategoryDTO{}.TableName()).
		Where("id IN ?", ids).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
