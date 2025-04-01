package models

type Coin struct {
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
}

// Our "fake database" — slice of coins
var Portfolio = []Coin{
	{Symbol: "BTC", Quantity: 0.1},
	{Symbol: "ETH", Quantity: 1.5},
}
