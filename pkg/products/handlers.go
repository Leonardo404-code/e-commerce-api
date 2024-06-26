package products

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
