package payment

import (
	"fmt"
	"net/http"

	"github.com/Leonargo404-code/e-commerce/internal/env"
	stripePkg "github.com/Leonargo404-code/e-commerce/pkg/stripe"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// @Summary		Payment indent
// @Description	Create a new payment indent in stripe
// @Tags			Stripe
// @Router			/products [post]
// @Param			data	body	stripe.Payment	true	"Required payment information"
// @Accept			json
// @Produce		json
// @Success		200	{object} string "{"client_secret": "0a52ffce-20ed-43d6-81eb-aa5deaf33f34"}"
func (h *handler) CreatePaymentIntent(c *fiber.Ctx) error {
	stripe.Key = env.GetString("STRIPE.TEST")

	var payment stripePkg.Payment

	if err := c.BodyParser(&payment); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": fmt.Errorf("error in body parser: %v", err),
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
			"error": fmt.Errorf("error in create payment intent: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"client_secret": pi.ClientSecret,
	})
}
