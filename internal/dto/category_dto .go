package dto

import "kasir-api/internal/domain"

type CategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1"`
	Description string `json:"description" validate:"max=255"`
}

func CategoryReqToDomain(req *CategoryRequest) *domain.Category {
	return &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
