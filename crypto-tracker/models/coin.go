package models

type Coin struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
	UserID   uint    `json:"user_id"`
}
