package dto

import "kasir-api/internal/model"

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func ProductReqToDomain(req *ProductRequest) *model.Product {
	return &model.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
}
