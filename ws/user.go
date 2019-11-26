package ws

import (
	"ele-dex/book"
	"ele-dex/market"
	"encoding/json"
	"log"
)

type User struct {
	client *Client
}

type T struct {
	Book   *book.Book     `json:"book"`
	Market *market.Market `json:"market"`
}

func (u *User) SetClient(client *Client) {
	u.client = client
}

func (u *User) Register() {
	book := book.TestBook()
	market := market.TestMarket()
	data := T{
		Book:   book,
		Market: market,
	}
	bytes, _ := json.Marshal(data)
	u.client.Send(bytes)
}

func (u *User) Unregister() {
	log.Println("user unregister")
}

func (u *User) ReceMsg(msg []byte) {
	log.Println("user rece msg", string(msg))
}
