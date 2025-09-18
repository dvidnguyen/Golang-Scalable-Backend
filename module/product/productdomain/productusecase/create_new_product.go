package productusecase

import (
	"Ls04_GORM/module/product/productdomain"
	"strings"

	"golang.org/x/net/context"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}

func NewCreateProductUseCase(repo CreateProductRepository) CreateNewProductUseCase {
	return CreateNewProductUseCase{
		repo: repo,
	}
}

type CreateNewProductUseCase struct {
	repo CreateProductRepository
}

func (uc CreateNewProductUseCase) CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error {
	prod.Name = strings.TrimSpace(prod.Name)

	if prod.Name == "" {
		return productdomain.ErrProductNameCannotBeBlank
	}

	if err := uc.repo.CreateProduct(ctx, prod); err != nil {
		return err
	}

	return nil
}

type CreateProductRepository interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}
