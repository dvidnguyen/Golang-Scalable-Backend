package productmysql

import (
	"Ls04_GORM/module/product/domain"
	"context"
)

func (repo MysqlRepository) CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error {
	if err := repo.db.Table(prod.TableName()).Create(&prod).Error; err != nil {
		return err
	}

	return nil
}
