package ws

import "log"

type User struct {
	client *Client
}

func (u *User) SetClient(client *Client) {
	u.client = client
}

func (u *User) Register() {
	log.Println("user register")
	u.client.Send([]byte("Hello client!"))
}

func (u *User) Unregister() {
	log.Println("user unregister")
}

func (u *User) ReceMsg(msg []byte) {
	log.Println("user rece msg", string(msg))
}
