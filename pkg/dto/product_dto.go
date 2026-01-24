package dto

import "kasir-api/pkg/domain"

type ProductRequest struct {
	Name  string `json:"name" validate:"required,min=1"`
	Price int    `json:"price" validate:"required,number"`
	Stock int    `json:"stock" validate:"required,number"`
}

func ProductReqToDomain(req *ProductRequest) *domain.Product {
	return &domain.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
}
