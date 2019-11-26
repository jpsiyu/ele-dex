package ws

import (
	"ele-dex/book"
	"encoding/json"
	"log"
)

type User struct {
	client *Client
}

func (u *User) SetClient(client *Client) {
	u.client = client
}

func (u *User) Register() {
	book := book.TestBook()
	bytes, _ := json.Marshal(book)
	u.client.Send(bytes)
}

func (u *User) Unregister() {
	log.Println("user unregister")
}

func (u *User) ReceMsg(msg []byte) {
	log.Println("user rece msg", string(msg))
}
