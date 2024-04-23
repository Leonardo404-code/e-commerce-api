package handler

import (
	"encoding/json"
	"net/http"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Create(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{})
	}

	productInfo := form.Value["product"][0]
	if len(productInfo) == 0 {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "error in get product info",
		})
	}

	productImage := form.File["image"][0]
	if productImage == nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "error in get product image",
		})
	}

	product := &productsPkg.Product{}
	if err := json.Unmarshal([]byte(productInfo), &product); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{})
	}

	newProduct, err := h.service.Create(product, productImage)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(newProduct)
}
