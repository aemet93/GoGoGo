package portfolio

type Coin struct {
	Name   string
	Amount float64
	Price  float64
}

func (c Coin) GetTotalUSD() float64 {
	return c.Amount * c.Price
}
