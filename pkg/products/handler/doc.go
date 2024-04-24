package handler

import (
	"time"

	"github.com/google/uuid"
)

type (
	ResultDoc struct {
		Products []*ProductDoc `json:"products,omitempty"`
		Total    int           `json:"total,omitempty"`
		Page     int           `json:"page,omitempty"`
	}

	ProductDoc struct {
		CreatedAt   time.Time `json:"created_at,omitempty" example:"1980-00-00T01:00:00.00000-03:00"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" example:"1980-00-00T01:00:00.00000-03:00"`
		DeletedAt   DeletedAt `json:"deleted_at,omitempty"`
		Name        string    `json:"name,omitempty" example:"product"`
		Description string    `json:"description,omitempty" example:"a nice product"`
		Image       string    `json:"image,omitempty" example:"http://storage.com/product"`
		Value       float64   `json:"value,omitempty" example:"500.75"`
		Units       int64     `json:"units,omitempty" example:"10"`
		ID          uuid.UUID `json:"id,omitempty" example:"a8fa2f1d-ca74-47cc-9631-28c89c8470a5"`
	}

	DeletedAt struct {
		Time  string `json:"Time" form:"Time"`
		Valid bool   `json:"Valid" form:"Valid"`
	}

	ProductReqDoc struct {
		Name        string  `json:"name,omitempty" example:"product"`
		Description string  `json:"description,omitempty" example:"a nice product"`
		Value       float64 `json:"value,omitempty" example:"500.75"`
		Units       int64   `json:"units,omitempty" example:"10"`
	}
)
