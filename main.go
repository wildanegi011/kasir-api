package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Product 1", Price: 10000, Stock: 10},
	{ID: 2, Name: "Product 2", Price: 20000, Stock: 20},
	{ID: 3, Name: "Product 3", Price: 30000, Stock: 30},
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  "ok",
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
func getProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(Response{
				Status:  "ok",
				Message: "succesfully get product",
				Data:    product,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found",
	})
}

// CreateProduct godoc
// @Summary Create new product
// @Description Menambahkan produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body Product true "Product Data"
// @Success 201 {object} map[string]interface{}
// @Router /api/products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Status:  "ok",
		Message: "successfully added product",
		Data:    products,
	})
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Product Data"
// @Success 200 {object} map[string]interface{}
// @Router /api/products/{id} [put]
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)
	var updateProduct Product
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range products {
		if products[i].ID == id {
			updateProduct.ID = id
			products[i] = updateProduct
			json.NewEncoder(w).Encode(Response{
				Status:  "ok",
				Message: "succesfully update product",
				Data:    products[i],
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found",
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
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(Response{
				Status:  "ok",
				Message: "succesfully delete product",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found",
	})
}

// @title Kasir API
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {

	http.HandleFunc("GET /api/products", GetProducts)
	http.HandleFunc("GET /api/products/", func(w http.ResponseWriter, r *http.Request) {
		getProductByID(w, r)
	})
	http.HandleFunc("PUT /api/products/", func(w http.ResponseWriter, r *http.Request) {
		updateProduct(w, r)
	})
	http.HandleFunc("DELETE /api/products/", func(w http.ResponseWriter, r *http.Request) {
		deleteProduct(w, r)
	})
	http.HandleFunc("POST /api/products", CreateProduct)

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Status:  "ok",
			Message: "server is running",
		})
	})

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("gagal running server")
	}
}
