package product

import "github.com/gofiber/fiber/v2"

func GetAllRouter(app fiber.Router, service Service) {
	app.Get("/products", GetAllHandler(service))
}

func GetByIDRouter(app fiber.Router, service Service) {
	app.Get("/products/:id", GetByIDHandler(service))
}

func CreateRouter(app fiber.Router, service Service) {
	app.Post("/products", CreateHandler(service))
}

func UpdateRouter(app fiber.Router, service Service) {
	app.Put("/products", UpdateHandler(service))
}

func DeleteRouter(app fiber.Router, service Service) {
	app.Delete("/products", DeleteHandler(service))
}

