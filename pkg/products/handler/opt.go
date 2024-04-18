package handler

import "github.com/Leonargo404-code/e-commerce/pkg/products"

type Option func(*handler) error

func WithService(productSvc products.Services) Option {
	return func(h *handler) error {
		h.service = productSvc
		return nil
	}
}
