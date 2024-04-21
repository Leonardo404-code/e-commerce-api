package handler

import (
	"net/http"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Get(c *fiber.Ctx) error {
	queries := c.Queries()
	filter, err := products.ParseFilters(queries)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	result, err := h.repo.Get(filter)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.JSON(result)
}
