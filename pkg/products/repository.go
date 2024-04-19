package products

type Repository interface {
	Get(page int) (*Result, error)
	Create(newProduct *Product) error
}
