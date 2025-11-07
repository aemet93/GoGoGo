package portfolio

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

func (p Portfolio) TotalInUSD() float64 {
	var total float64
	for _, coin := range p.Coins {
		total += coin.GetTotalUSD()
	}

	return total
}

func (p Portfolio) Buy(name string, amount float64) {
	coin := p.Coins[name]
	coin.Amount += amount
	p.Coins[name] = coin
}
