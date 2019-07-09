package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string
	emai string
}

func (u *user) notify() {
	fmt.Printf("Send mail for user: %s<%s>\n",
		u.name,
		u.emai)
}

type admin struct {
	name string
	emai string
}

func (a *admin) notify() {
	fmt.Printf("Send mail for admin: %s<%s>\n",
		a.name,
		a.emai)
}

func main() {
	u := user{"Logan", "hs.lee.login@gmail.com"}
	// pointer, becouse notify() use pointer receiver
	sendNotification(&u)

	a := admin{"Hyunsuk", "blackknight8012@gmail.com"}
	// pointer, becouse notify() use pointer receiver
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
