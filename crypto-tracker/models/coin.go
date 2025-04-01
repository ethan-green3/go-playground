package models

type Coin struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
	UserID   uint    `json:"user_id"`
}

// Our "fake database" — slice of coins
var Portfolio = []Coin{
	{Symbol: "BTC", Quantity: 0.1},
	{Symbol: "ETH", Quantity: 1.5},
}
