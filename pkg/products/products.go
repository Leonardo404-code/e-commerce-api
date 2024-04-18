package products

import "gorm.io/gorm"

type (
	Product struct {
		gorm.Model
		Name        string  `gorm:"size:255; not null;" json:"name"`
		Description string  `gorm:"null" json:"description"`
		Value       float64 `gorm:"not null;" json:"value"`
		Units       int64   `gorm:"not null" json:"units"`
		Image       Image   `gorm:"foreingKey:ProductID; references:ID; constraint:OnDelete:SET NULL;" json:"photo"`
	}

	Image struct {
		ProductID uint64 `gorm:"not null" json:"product_id"`
		Url       string `gorm:"not null;" json:"url"`
		FileName  string `gorm:"not null" json:"file_name"`
	}

	Result struct {
		Products []Product `json:"products,omitempty"`
		Total    int       `json:"total,omitempty"`
		Page     int       `json:"page,omitempty"`
	}
)
