package main

import (
	"log"

	"github.com/aliyilmazdev/go-simple-rest-api/internal/product"
	"github.com/aliyilmazdev/go-simple-rest-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	db := database.Connect()

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	product.CreateRouter(app, productService)
	product.GetByIDRouter(app, productService)
	product.GetAllRouter(app, productService)
	product.UpdateRouter(app, productService)
	product.DeleteRouter(app, productService)

	log.Fatal(app.Listen(":3000")) 
}