package controller

import (
	"Ls04_GORM/module/product/productdomain"
	"context"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewAPIController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase: createUseCase}
}
