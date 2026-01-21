package dto

import "kasir-api/internal/domain"

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func ProductReqToDomain(req *ProductRequest) *domain.Product {
	return &domain.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
}
