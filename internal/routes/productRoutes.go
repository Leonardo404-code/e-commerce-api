package routes

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products/handler"
	"github.com/Leonargo404-code/e-commerce/pkg/products/service"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	handlers := handler.Must(
		service.Must(),
	)

	app.Get("/products", handlers.Get)
	app.Post("/products", handlers.Create)
}
