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
	user  // Embedded type.
	level string
}

func main() {
	a := admin{
		user:  user{"Logan", "hs.lee.login@gmail.com"},
		level: "super",
	}
	// Access inner type method directly.
	a.user.notify()

	// Inner type method is promoted.
	a.notify()

	// Interface is promoted.
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
