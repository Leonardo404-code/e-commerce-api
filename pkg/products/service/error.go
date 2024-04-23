package service

import "errors"

var (
	ErrGetProduct       = errors.New("error in get products")
	ErrProductsNotFound = errors.New("error in found products")
	ErrProductsConflict = errors.New("conflict occurred in the search for products")
	ErrDeleteProduct    = errors.New("error in delete products")
)
