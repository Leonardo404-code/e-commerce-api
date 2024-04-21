package products

type Services interface {
	Create(*Product) (*Product, error)
	Delete(*Filter) error
}
