package user

import "log"

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) Register() {
	log.Println("user register")
}

func (u *User) Unregister() {
	log.Println("user unregister")
}

func (u *User) ReceMsg(msg []byte) {
	log.Println("user rece msg", string(msg))
}
