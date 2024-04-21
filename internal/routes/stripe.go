package routes

import (
	"github.com/Leonargo404-code/e-commerce/pkg/stripe/payment"
	"github.com/gofiber/fiber/v2"
)

func StripeRoutes(app *fiber.App) {
	handlers := payment.Must()

	app.Post("/pay", handlers.CreatePaymentIntent)
}
