package products

import (
	"database/sql"
	"time"

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
