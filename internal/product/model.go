package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string
	Price float64
}

type UpdateProductRequest struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type DeleteProductRequest struct {
	Id uint `json:"id"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type GetAllProductResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}