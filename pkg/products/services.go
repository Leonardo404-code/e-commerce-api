package products

import "mime/multipart"

type Services interface {
	Create(product *Product, productImage *multipart.FileHeader) (*Product, error)
	Delete(filter *Filter) error
}
