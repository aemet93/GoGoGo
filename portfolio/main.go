package portfolio

func main() {
	var portfolio Portfolio
	_ = portfolio

	var coins map[string]Coin
	coins = map[string]Coin{
		"ETH": Coin{
			Name:   "Ethereum",
			Amount: 0,
			Price:  4120,
		},
		"BTC": Coin{
			Name:   "Bitcoin",
			Amount: 0,
			Price:  115200,
		},
	}

	portfolio = Portfolio{
		Coins: coins,
	}
}
