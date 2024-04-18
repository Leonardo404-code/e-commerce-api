package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (h *handler) Get(c fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error":   err,
			"message": "error to get the query page",
		})
	}

	result, err := h.service.Get(page)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.JSON(result)
}
