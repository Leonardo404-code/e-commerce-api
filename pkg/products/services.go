package products

type Services interface {
	Get(*Filter) (*Result, error)
	Create(product *Product) (*Product, error)
}
