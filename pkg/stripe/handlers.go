package stripe

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreatePaymentIntent(c *fiber.Ctx) error
}
