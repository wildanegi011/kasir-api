package dto

import "kasir-api/pkg/domain"

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CategoryReqToDomain(req *CategoryRequest) *domain.Category {
	return &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
