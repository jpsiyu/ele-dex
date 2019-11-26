package book

import (
	"math/rand"
)

type Order struct {
	Price       float64 `json:"price"`
	BaseCoin    string  `json:"baseCoin"`
	TradeCoin   string  `json:"tradeCoin"`
	BaseAmount  float64 `json:"baseAmount"`
	TradeAmount float64 `json:"tradeAmount"`
}

type Book struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

func TestBook() *Book {
	book := Book{
		Asks: randomOrderList(),
		Bids: randomOrderList(),
	}
	return &book
}

func randomOrderList() []Order {
	orderList := []Order{}
	for i := 0; i < 10; i++ {
		order := Order{
			Price:       rand.Float64(),
			BaseCoin:    "ETH",
			TradeCoin:   "IDEX",
			BaseAmount:  rand.Float64(),
			TradeAmount: rand.Float64(),
		}
		orderList = append(orderList, order)
	}
	return orderList
}
