package products

type Repository interface {
	Get(*Filter) (*Result, error)
	Create(newProduct *Product) error
}
