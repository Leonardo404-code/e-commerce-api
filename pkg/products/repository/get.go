package repository

import (
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (r *repository) Get(filter *productsPkg.Filter) (*productsPkg.Result, error) {
	var products []productsPkg.Product

	totalPages := r.db.Find(&products).RowsAffected
	if totalPages == 0 {
		totalPages = 1
	}

	if len(filter.ID) != 0 {
		err := r.db.Where("id IN ?", filter.ID).Offset((filter.Page - 1) * 10).Limit(10).Order("id desc").Find(&products)
		if err.Error != nil {
			return nil, fmt.Errorf("failed in find products: %v", err.Error)
		}
	} else {
		err := r.db.Offset((filter.Page - 1) * 10).Limit(10).Order("id desc").Find(&products)
		if err.Error != nil {
			return nil, fmt.Errorf("failed in find products: %v", err.Error)
		}
	}

	return &productsPkg.Result{
		Products: products,
		Total:    int(totalPages),
		Page:     filter.Page,
	}, nil
}
