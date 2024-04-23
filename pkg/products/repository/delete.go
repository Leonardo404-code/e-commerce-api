package repository

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"gorm.io/gorm"
)

func (r *repository) Delete(filter *products.Filter, deleteFromBucket func() error) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&products.Product{}, filter.ID); err.Error != nil {
			return err.Error
		}

		if err := deleteFromBucket(); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
