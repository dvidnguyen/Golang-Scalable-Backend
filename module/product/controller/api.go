package controller

import (
	"Ls04_GORM/module/product/domain"
	"context"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewAPIController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase: createUseCase}
}
