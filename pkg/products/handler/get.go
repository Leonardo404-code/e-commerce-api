package handler

import (
	"net/http"
	"strconv"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Get(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error":   err,
			"message": "error to get the query page",
		})
	}

	productId := c.Queries()
	filter := products.ParseFilters(page, productId)

	result, err := h.service.Get(filter)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.JSON(result)
}
