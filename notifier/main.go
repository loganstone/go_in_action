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

func main() {
	u := user{"Logan", "hs.lee.login@gmail.com"}
	// pointer, becouse notify() use pointer receiver
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
