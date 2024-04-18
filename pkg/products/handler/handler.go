package handler

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

type handler struct{}

func new() (products.Handlers, error) {
	h := &handler{}

	return h, nil
}

func Must() products.Handlers {
	p, err := new()
	if err != nil {
		panic(err)
	}

	return p
}
