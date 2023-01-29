package types

type PriceResponse struct {
	Ticket string  `json:"ticket"`
	Price  float64 `json:"price"`
}
