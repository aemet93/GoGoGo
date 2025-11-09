package portfolio

import "log"

func RunPortfolio() {
	var portfolio Portfolio

	var coins map[string]Coin
	coins = map[string]Coin{
		"ETH": Coin{
			Name:  "Ethereum",
			Price: 4120,
		},
		"BTC": Coin{
			Name:  "Bitcoin",
			Price: 115200,
		},
	}

	portfolio = Portfolio{
		Coins: coins,
	}

	log.Printf("Total coins %v", portfolio.CoinCount())

	portfolio.Buy("ETH", 2)

	portfolio.Buy("BTC", 10)

	log.Printf("Total value %v USD", portfolio.TotalInUSD())

	portfolio.Sell("BTC", 10)

	portfolio.Sell("ETH", 0.75)

	log.Printf("Total value %v USD", portfolio.TotalInUSD())

	log.Printf("ETH balance %v", portfolio.Balance("ETH"))

	log.Printf("BTC balance %v", portfolio.Balance("BTC"))
}
