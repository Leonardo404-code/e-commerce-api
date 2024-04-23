package routes

import (
	"github.com/Leonargo404-code/e-commerce/pkg/google"
	"github.com/Leonargo404-code/e-commerce/pkg/products/handler"
	"github.com/Leonargo404-code/e-commerce/pkg/products/repository"
	"github.com/Leonargo404-code/e-commerce/pkg/products/service"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	repository := repository.Must()
	google := google.Must()
	handlers := handler.Must(
		service.Must(repository, google),
		repository,
	)

	app.Get("/products", handlers.Get)
	app.Post("/products", handlers.Create)
	app.Delete("/products", handlers.Delete)
}
