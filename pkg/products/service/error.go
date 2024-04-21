package service

import "errors"

var (
	ErrGetProduct       = errors.New("error in get products")
	ErrProductsNotFound = errors.New("error in found products")
	ErrDeleteProduct    = errors.New("error in delete products")
)
