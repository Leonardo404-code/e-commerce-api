package products

import "github.com/gofiber/fiber/v3"

type Handlers interface {
	Get(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
	UploadPhoto(c fiber.Ctx) error
}
