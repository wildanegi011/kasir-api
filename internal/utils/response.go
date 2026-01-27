package utils

import (
	"encoding/json"
	"math"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Metadata struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalPage int `json:"total_pages"`
}

type MetadataOption func(*Metadata)

func WithPagination(total, page, pageSize int) MetadataOption {
	return func(m *Metadata) {
		m.Total = total
		m.Page = page
		m.PageSize = pageSize
		m.TotalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	}
}

type ResponseWithMetadata struct {
	Response
	Metadata *Metadata `json:"metadata,omitempty"`
}

func SuccessResponse(
	w http.ResponseWriter,
	statusCode int,
	message string,
	data interface{},
	opts ...MetadataOption,
) error {

	var meta *Metadata
	if len(opts) > 0 {
		m := &Metadata{}
		for _, opt := range opts {
			opt(m)
		}
		meta = m
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if meta != nil {
		return json.NewEncoder(w).Encode(ResponseWithMetadata{
			Response: Response{
				Status:  true,
				Message: message,
				Data:    data,
			},
			Metadata: meta,
		})
	}

	return json.NewEncoder(w).Encode(Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(
	w http.ResponseWriter,
	statusCode int,
	message string,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(Response{
		Status:  false,
		Message: message,
	})
}
