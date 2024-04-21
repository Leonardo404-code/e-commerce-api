package payment

import (
	"net/http"

	"github.com/Leonargo404-code/e-commerce/internal/env"
	stripePkg "github.com/Leonargo404-code/e-commerce/pkg/stripe"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func (h *handler) CreatePaymentIntent(c *fiber.Ctx) error {
	stripe.Key = env.GetString("STRIPE_TEST")

	var payment stripePkg.Payment

	if err := c.BodyParser(&payment); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "error in body-parser",
		})
	}

	params := &stripe.PaymentIntentParams{
		Amount:   &payment.Amount,
		Currency: stripe.String(string(stripe.CurrencyBRL)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		ReceiptEmail: &payment.Email,
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"client_secret": pi.ClientSecret,
	})
}
