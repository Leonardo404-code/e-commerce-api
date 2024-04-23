package repository

import (
	"fmt"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"gorm.io/gorm"
)

func (r *repository) Create(newProduct *products.Product, upload func() error) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newProduct).Error; err != nil {
			return fmt.Errorf("failed to create a new product: %v", err)
		}

		if err := upload(); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
