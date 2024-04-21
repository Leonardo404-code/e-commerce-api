package stripe

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Amount    int64        `gorm:"not null;" json:"amount"`
	Email     string       `gorm:"not null;" json:"email"`
	ID        uuid.UUID    `gorm:"primarykey"`
}
