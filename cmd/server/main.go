package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir-api/internal/config"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"
	"kasir-api/internal/utils"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {

	db, _ := config.InitDB()

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("GET /api/products", productHandler.GetProducts)
	http.HandleFunc("GET /api/products/", productHandler.GetProductByID)
	http.HandleFunc("POST /api/products", productHandler.CreateProduct)
	http.HandleFunc("PUT /api/products/", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/", productHandler.DeleteProduct)

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  true,
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
