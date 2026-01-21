package handler

import (
	"encoding/json"
	"errors"
	"kasir-api/internal/dto"
	"kasir-api/internal/service"
	"kasir-api/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := h.productService.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to get products",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully get products",
		Data:    products,
	})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid Request",
		})
		return
	}
	product, err := h.productService.GetProductByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to get product",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully get product",
		Data:    product,
	})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Membuat produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductRequest true "Product Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.Price <= 0 || req.Stock < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Name, price (must be > 0), and stock (must be >= 0) are required",
		})
		return
	}

	product := dto.ProductReqToDomain(&req)

	if _, err := h.productService.CreateProduct(product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Failed to create product",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "Product created successfully",
		Data:    product,
	})
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.ProductRequest true "Product Data"
// @Success 200 {object} map[string]interface{}
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, _ := strconv.Atoi(id)

	var req dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.Price <= 0 || req.Stock < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Name, price (must be > 0), and stock (must be >= 0) are required",
		})
		return
	}

	product := dto.ProductReqToDomain(&req)
	updatedProduct, err := h.productService.UpdateProduct(idInt, product)
	if err != nil {
		if errors.Is(err, utils.ErrProductNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.Response{
				Status:  false,
				Message: utils.ErrProductNotFound.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Failed to update product",
		})
		return
	}

	updatedProduct.ID = idInt

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "Product updated successfully",
		Data:    updatedProduct,
	})
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid Request",
		})
		return
	}
	if err := h.productService.DeleteProduct(idInt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to delete product",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully delete product",
	})
}
