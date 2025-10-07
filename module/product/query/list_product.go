package query

import (
	"Ls04_GORM/common"
	"context"

	"github.com/google/uuid"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type ProductDTO struct {
	common.BaseModel
	CategoryId  uuid.UUID    `gorm:"column:category_id" json:"category_id"`
	Name        string       `gorm:"column:name" json:"name"`
	Type        string       `gorm:"column:type" json:"type"`
	Description string       `gorm:"column:description" json:"description"`
	Category    *CategoryDTO `gorm:"foreignKey:CategoryId;references:Id" json:"category"`
}

type CategoryDTO struct {
	Id    uuid.UUID `gorm:"column:id" json:"id"`
	Title string    `gorm:"column:title" json:"title"`
}

func (CategoryDTO) TableName() string {
	return "categories"
}

type ListProductFilter struct {
	CategoryId string `form:"category_id" json:"category_id"`
}

type listProductQuery struct {
	sctx sctx.ServiceContext
}

func NewListProductQuery(sctx sctx.ServiceContext) listProductQuery {
	return listProductQuery{sctx: sctx}
}

type ListProductParam struct {
	Filter ListProductFilter
	common.Paging
}

func (p listProductQuery) Execute(ctx context.Context, param *ListProductParam) ([]ProductDTO, error) {
	var products []ProductDTO
	dbContext := p.sctx.MustGet(common.KeyGorm).(common.DBContext)

	db := dbContext.GetDB().Table("products")

	if param.Filter.CategoryId != "" {
		db = db.Where("category_id = ?", param.Filter.CategoryId)
	}

	var count int64
	db.Count(&count)
	param.Total = int(count)

	param.Process()

	offset := param.Limit * (param.Page - 1)

	if err := db.Offset(offset).Limit(param.Limit).Order("id desc").Find(&products).Error; err != nil {
		return nil, core.ErrBadRequest.WithError("Cannot list product").WithDebug(err.Error())
	}

	return products, nil

}
