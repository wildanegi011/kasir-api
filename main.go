package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

func getProductByID(id int, w http.ResponseWriter) {
	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "ok",
				"message": "succesfully get product",
				"data":    product,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "error",
		"message": "product not found",
	})
}
func updateProduct(id int, w http.ResponseWriter, r *http.Request) {
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
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "ok",
				"message": "succesfully update product",
				"data":    products[i],
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "error",
		"message": "product not found",
	})

}

func deleteProduct(id int, w http.ResponseWriter) {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "ok",
				"message": "succesfully delete product",
				"data":    products,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "error",
		"message": "product not found",
	})
}

func main() {
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case "GET":
			getProductByID(id, w)
		case "PUT":
			updateProduct(id, w, r)
		case "DELETE":
			deleteProduct(id, w)
		}
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "ok",
				"message": "succesfully get product",
				"data":    products,
			})
		case "POST":
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			newProduct.ID = len(products) + 1
			products = append(products, newProduct)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "ok",
				"message": "successfully added product",
				"data":    products,
			})
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "server is running",
		})
	})
	fmt.Println("server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("gagal running server")
	}
}
