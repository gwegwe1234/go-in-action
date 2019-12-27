package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("어드민에게 메세지를 전달합니다. : %s<%s>\n", a.name, a.email)
}

func (u *user) notify() {
	fmt.Printf("사용자에게 메세지를 전달합니다. : %s<%s>\n", u.name, u.email)
}

func main() {
	jade := user{"jade.jjj", "jade.jjj@kikiki.com"}
	judy := admin{"judy.jucie", "judy.juice@qwert.com"}

	sendNotification(&jade)
	sendNotification(&judy)
}

func sendNotification(n notifier) {
	n.notify()
}
