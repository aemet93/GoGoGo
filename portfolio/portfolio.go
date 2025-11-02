package portfolio

type Coin struct {
	Name   string
	Amount float64
	Price  float64
}

func (c Coin) TotalInUSD() float64 {
	return c.Amount * c.Price
}

type Portfolio struct {
	Coins map[string]Coin
}

func (p Portfolio) GetCoins() map[string]float64 {
	coinList := map[string]float64{}

	for _, coin := range p.Coins {
		coinList[coin.Name] = coin.Amount
	}

	return coinList
}
