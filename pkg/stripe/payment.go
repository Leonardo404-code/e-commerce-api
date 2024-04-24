package stripe

type Payment struct {
	Amount int64  `gorm:"not null;" json:"amount" example:"500"`
	Email  string `gorm:"not null;" json:"email" example:"leonardobispo.dev@gmail.com"`
}
