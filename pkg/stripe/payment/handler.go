package payment

import "github.com/Leonargo404-code/e-commerce/pkg/stripe"

type handler struct{}

func Must() stripe.Handler {
	return &handler{}
}
