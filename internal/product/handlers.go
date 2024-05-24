package product

import (
	"fmt"
	"strconv"

	"github.com/aliyilmazdev/go-simple-rest-api/pkg/presenter"
	"github.com/gofiber/fiber/v2"
)

func GetAllHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := service.getAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(products))
	}
}

func GetByIDHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
		}
		product, err := service.getByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(product))
	}
}

func CreateHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createProductRequest CreateProductRequest
		if err := c.BodyParser(&createProductRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
		}
		if err := service.create(createProductRequest); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(nil))
	}
}

func UpdateHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var updateProductRequest UpdateProductRequest
		id := c.Params("id")

		   if id == "" {
            return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(fmt.Errorf("id parameter is missing")))
        }
		
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
		}
		if err := c.BodyParser(&updateProductRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
		}
		if err := service.update(uint(idUint), updateProductRequest); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(nil))
	}
}

func DeleteHandler(service Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var deleteProductRequest DeleteProductRequest
		if err := c.BodyParser(&deleteProductRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(err))
		}
		if err := service.delete(int(deleteProductRequest.Id)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(nil))
	}
}