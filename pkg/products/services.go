package products

type Services interface {
	Create(product *Product) (*Product, error)
}
