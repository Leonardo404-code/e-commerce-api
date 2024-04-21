package products

import "strings"

type Filter struct {
	ID   []string
	Page int
}

func ParseFilters(page int, productID map[string]string) *Filter {
	var productsIDs []string
	if productID["id"] != "" {
		productsIDs = strings.Split(productID["id"], ",")
	}

	return &Filter{
		ID:   productsIDs,
		Page: page,
	}
}
