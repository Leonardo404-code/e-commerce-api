package handler

import (
	"net/http"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Create(c *fiber.Ctx) error {
	product := &productsPkg.Product{}

	if err := c.BodyParser(product); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "failed to parse-body",
		})
	}

	newProduct, err := h.service.Create(product)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(newProduct)
}
