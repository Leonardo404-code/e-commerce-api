package products

type Services interface {
	Get(page int) (*Result, error)
}
