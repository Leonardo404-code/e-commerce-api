package products

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Leonargo404-code/e-commerce/internal/env"
	"github.com/google/uuid"
)

type (
	Product struct {
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   sql.NullTime `gorm:"index"`
		Name        string       `gorm:"size:255; not null;" json:"name"`
		Description string       `gorm:"null" json:"description"`
		Image       string       `gorm:"null" json:"url"`
		Value       float64      `gorm:"not null;" json:"value"`
		Units       int64        `gorm:"not null" json:"units"`
		ID          uuid.UUID    `gorm:"primarykey"`
	}

	Result struct {
		Products []Product `json:"products,omitempty"`
		Total    int       `json:"total,omitempty"`
		Page     int       `json:"page,omitempty"`
	}
)

func (p *Product) FormatName() string {
	productNameTrim := strings.Trim(p.Name, " ")
	productNameToLower := strings.ToLower(productNameTrim)
	return productNameToLower
}

func (p *Product) FormatStorageURL(StorageURL, BucketName string) string {
	productNameTrim := strings.Trim(p.Name, " ")
	countSpaces := strings.Count(productNameTrim, " ")
	productNameReplace := strings.Replace(productNameTrim, " ", "-", countSpaces)
	productNameToLower := strings.ToLower(productNameReplace)

	imageURL := fmt.Sprintf(
		"%s/%s/%s",
		env.GetString(StorageURL),
		env.GetString(BucketName),
		productNameToLower,
	)

	return imageURL
}
