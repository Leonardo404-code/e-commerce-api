package handler

import (
	"net/http"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get Products
// @Description	Get products matching passed filters
// @Tags			Products
// @Router			/products [get]
// @Param			id			query	string	false	"Products ID you want"
// @Param			page		query	int		true	"Current page for pagination system"
// @Param			maxPerPage query	int		true	"Maximum number of items per page"
// @Produce		json
// @Success		200	{object}	handler.ResultDoc
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
