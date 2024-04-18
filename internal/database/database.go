package database

import (
	"fmt"

	"github.com/Leonargo404-code/e-commerce/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: env.GetString("DATABASEURL"),
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed in database connection: %v", err))
	}

	return db
}
