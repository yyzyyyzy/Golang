package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"LZK", "916990143@qq.com"}
	sendNotification(&u)
}
