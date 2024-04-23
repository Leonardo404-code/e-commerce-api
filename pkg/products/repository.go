package products

type Repository interface {
	Get(filter *Filter) (*Result, error)
	Create(newProduct *Product, upload func() error) error
	Delete(filter *Filter) error
}
