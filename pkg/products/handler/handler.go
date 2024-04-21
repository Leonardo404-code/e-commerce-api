package handler

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

type handler struct {
	service products.Services
	repo    products.Repository
}

func new(opts ...Option) (products.Handlers, error) {
	h := &handler{}

	for _, opt := range opts {
		if err := opt(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

func Must(productSvc products.Services, productsRepo products.Repository) products.Handlers {
	p, err := new(
		WithService(productSvc),
	)
	if err != nil {
		panic(err)
	}

	return p
}
