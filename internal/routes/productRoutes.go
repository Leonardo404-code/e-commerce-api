package routes

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products/handler"
	"github.com/gofiber/fiber/v3"
)

func ProductRoutes(app *fiber.App) {
	handlers := handler.Must()

	app.Get("/products", handlers.Get)
}
