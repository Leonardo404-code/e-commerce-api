package handler

import (
	"net/http"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Delete Products
// @Description	Delete products from database matching passed filters
// @Tags			Products
// @Router			/products [delete]
// @Param			id			query	string	false	"Products ID you want to delete"
// @Param			page		query	int		true	"Current page for pagination system"
// @Param			maxPerPage	query	int		true	"Maximum number of items per page"
// @Success		200			"OK"
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
