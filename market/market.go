package market

import "math/rand"

type CoinPair struct {
	Coin  string  `json:"coin"`
	Star  bool    `json:"star"`
	Price float64 `json:"price"`
	Vol   float64 `json:"vol"`
	Chg   float64 `json:"chg"`
	Name  string  `json:"name"`
}

type CoinPairGroup struct {
	Base  string     `json:"base"`
	Pairs []CoinPair `json:"pairs"`
}

type Market struct {
	Groups []CoinPairGroup `json:"groups"`
}

func TestMarket() *Market {
	groups := []CoinPairGroup{
		CoinPairGroup{
			Base:  "ETH",
			Pairs: randomPair(),
		},
		CoinPairGroup{
			Base:  "TUSD",
			Pairs: randomPair(),
		},
		CoinPairGroup{
			Base:  "USDC",
			Pairs: randomPair(),
		},
		CoinPairGroup{
			Base:  "EURS",
			Pairs: randomPair(),
		},
		CoinPairGroup{
			Base:  "DAI",
			Pairs: randomPair(),
		},
	}

	return &Market{Groups: groups}
}

func randomPair() []CoinPair {
	pairs := []CoinPair{}
	for i := 0; i < 20; i++ {
		pair := CoinPair{
			Coin:  "ONT",
			Star:  false,
			Price: rand.Float64(),
			Vol:   rand.Float64(),
			Chg:   rand.Float64(),
			Name:  "Quant",
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
