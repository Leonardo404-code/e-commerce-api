package handler

import (
	"net/http"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Delete(c *fiber.Ctx) error {
	filter, err := productsPkg.ParseFilters(c.Queries())
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	if err := h.service.Delete(filter); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.JSON(http.StatusOK)
}
