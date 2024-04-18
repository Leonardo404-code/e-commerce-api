package repository

import (
	"github.com/Leonargo404-code/e-commerce/internal/database"
	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func Must() products.Repository {
	return repository{
		db: database.Connect(),
	}
}
