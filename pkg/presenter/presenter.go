package presenter

import "github.com/gofiber/fiber/v2"

func SuccessResponse(response interface{}) *fiber.Map {
	return &fiber.Map{
		"status":  true,
		"data":    response,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": err.Error(),
	}
}