# KASIR API (GOLANG)

## Description
Ini adalah aplikasi kasir sederhana yang dibuat menggunakan bahasa pemrograman golang

## Fitur
- Get all products
- Get product by id
- Update product by id
- Delete product by id
- Create product

## Instalasi Swagger
Download Swag
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
generate required files(docs folder and docs/doc.go).
```bash
swag init
```
Download http-swagger:
```bash
go get -u github.com/swaggo/http-swagger
```

## Cara menjalankan
```bash
go run main.go
```

## Endpoint API
- `GET /products` - Get all products
- `GET /products/:id` - Get product by id
- `PUT /products/:id` - Update product by id
- `DELETE /products/:id` - Delete product by id
- `POST /products` - Create product

## 1. Package dan Import
```go
package main
```
- Package utama yang berisi fungsi `main()`
- File ini bisa langsung dijalankan (go run main.go)

```go
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)
```
- `encoding/json` - untuk encoding dan decoding json
- `fmt` - untuk print ke terminal
- `net/http` - untuk membuat server dan handling request
- `strconv` - untuk konversi string ke integer
- `strings` - untuk handling string

## 2. Struct Response
```go
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
```
Struct ini digunakan untuk mengembalikan response dalam format JSON.

## 3. Struct Product
```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
```
struct adalah tipe data yang digunakan untuk mengelompokkan beberapa nilai. Sama seperti object di javascript atau object di bahasa pemrograman lain.

- `ID` - ID product
- `Name` - Nama product
- `Price` - Harga product
- `Stock` - Stok product

Tag `json:"..."` adalah tag yang digunakan untuk menentukan nama field dalam JSON. Misalnya `json:"id"` akan menghasilkan `"id"` dalam JSON.

## 4. Data Product (Dummy)
```go
var products = []Product{
	{ID: 1, Name: "Product 1", Price: 10000, Stock: 10},
	{ID: 2, Name: "Product 2", Price: 20000, Stock: 20},
	{ID: 3, Name: "Product 3", Price: 30000, Stock: 30},
}
```
Data product (dummy) yang digunakan untuk testing API.
- Ini adalah database sementara (in-memory) yang digunakan untuk testing API.
- Disimpan dalam memory (RAM) dan akan hilang ketika server dimatikan.

- `[]Product{}` adalah slice dari struct Product yang digunakan untuk menyimpan data product.
- Berikut ini adalah contoh perbedaan slice dengan array, slice bisa diubah ukurannya, sedangkan array tidak bisa:

```go
// Array (ukuran fixed)
arr := [3]int{1, 2, 3}
// arr[3] = 4 // ERROR! ukuran fixed

// Slice (ukuran bisa berubah)
slice := []int{1, 2, 3}
slice = append(slice, 4) // OK! ukuran bisa bertambah
```


## 5. Get All Products
```go
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  "ok",
		Message: "successfully get products",
		Data:    products,
	})
}
```
Fungsi ini digunakan untuk mengambil semua data product yang tersimpan dalam variabel `products`.
- `w` adalah response writer yang digunakan untuk mengirimkan response ke client.
- `r` adalah request yang diterima dari client.
- `w.Header().Set("Content-Type", "application/json")` digunakan untuk menentukan tipe content yang dikirimkan ke client.
- `json.NewEncoder(w).Encode(Response{})` digunakan untuk mengencode data product ke format JSON dan mengirimkannya ke client.

## 5. Get Product by ID
```go
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
```
Fungsi ini digunakan untuk mengambil data product berdasarkan ID yang diterima dari client.
- `w` adalah response writer yang digunakan untuk mengirimkan response ke client.
- `r` adalah request yang diterima dari client.
- `strings.TrimPrefix(r.URL.Path, "/products/")` digunakan untuk mengambil ID dari path URL.
- `strconv.Atoi(id)` digunakan untuk mengkonversi ID dari string ke integer.
- `for _, product := range products` digunakan untuk mencari product yang sesuai dengan ID.
- `w.WriteHeader(http.StatusNotFound)` digunakan untuk menentukan status code 404 jika product tidak ditemukan.
- `json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found"})` digunakan untuk mengencode error message ke format JSON dan mengirimkannya ke client.

## 6. Create Product
```go
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
```
Fungsi ini digunakan untuk membuat data product baru.
- `w` adalah response writer yang digunakan untuk mengirimkan response ke client.
- `r` adalah request yang diterima dari client.
- `json.NewDecoder(r.Body).Decode(&product)` digunakan untuk mengdecode data product dari request body.
- `products = append(products, product)` digunakan untuk menambahkan product baru ke slice `products`.
- `json.NewEncoder(w).Encode(product)` digunakan untuk mengencode data product yang baru dibuat ke format JSON dan mengirimkannya ke client.

## 7. Update Product
```go
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
```
Fungsi ini digunakan untuk mengupdate data product berdasarkan ID yang diterima dari client.
- `w` adalah response writer yang digunakan untuk mengirimkan response ke client.
- `r` adalah request yang diterima dari client.
- `strings.TrimPrefix(r.URL.Path, "/products/")` digunakan untuk mengambil ID dari path URL.
- `strconv.Atoi(id)` digunakan untuk mengkonversi ID dari string ke integer.
- `for i, product := range products` digunakan untuk mencari product yang sesuai dengan ID.
- `json.NewDecoder(r.Body).Decode(&products[i])` digunakan untuk mengdecode data product dari request body.
- `json.NewEncoder(w).Encode(products[i])` digunakan untuk mengencode data product yang diupdate ke format JSON dan mengirimkannya ke client.
- `w.WriteHeader(http.StatusNotFound)` digunakan untuk menentukan status code 404 jika product tidak ditemukan.
- `json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found"})` digunakan untuk mengencode error message ke format JSON dan mengirimkannya ke client.

## 8. Delete Product
```go
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

```
Fungsi ini digunakan untuk menghapus data product berdasarkan ID yang diterima dari client.
- `w` adalah response writer yang digunakan untuk mengirimkan response ke client.
- `r` adalah request yang diterima dari client.
- `strings.TrimPrefix(r.URL.Path, "/products/")` digunakan untuk mengambil ID dari path URL.
- `strconv.Atoi(id)` digunakan untuk mengkonversi ID dari string ke integer.
- `for i, product := range products` digunakan untuk mencari product yang sesuai dengan ID.
- `products = append(products[:i], products[i+1:]...)` digunakan untuk menghapus product dari slice `products`.
- `json.NewEncoder(w).Encode(Response{
				Status:  "ok",
				Message: "product deleted",
			})` digunakan untuk mengencode message ke format JSON dan mengirimkannya ke client.
- `w.WriteHeader(http.StatusNotFound)` digunakan untuk menentukan status code 404 jika product tidak ditemukan.
- `json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "product not found",
	})` digunakan untuk mengencode error message ke format JSON dan mengirimkannya ke client.


## License
MIT
