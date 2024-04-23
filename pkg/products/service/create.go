package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/google/uuid"
)

// TODO: ADD URL TO NEW PRODUCT STRUCT
func (s *service) Create(
	product *productsPkg.Product,
	productImage *multipart.FileHeader,
) (*productsPkg.Product, error) {
	if product.Name == "" {
		return nil, fmt.Errorf("name fild canot be empty")
	}

	if product.Value == 0.0 {
		return nil, fmt.Errorf("value cannot be zero")
	}

	if product.Units == 0 {
		return nil, fmt.Errorf("you need to have more than one unit to create the product")
	}

	imageURL := product.FormatStorageURL(STORAGEURL, BUCKETNAME)

	newProduct := &productsPkg.Product{
		ID:          uuid.New(),
		Name:        product.FormatName(),
		Description: product.Description,
		Value:       product.Value,
		Units:       product.Units,
		Image:       imageURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	imageFile, err := productImage.Open()
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(newProduct, s.uploadToBucket(newProduct.Name, imageFile)); err != nil {
		return nil, fmt.Errorf("error in create product: %v", err)
	}

	return newProduct, nil
}

func (s *service) uploadToBucket(
	productName string,
	productImage multipart.File,
) func() error {
	return func() error {
		return s.storage.Upload(context.Background(), productName, productImage)
	}
}
