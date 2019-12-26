package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify()  {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n", u.name, u.email)
}

func main()  {
	jinny := user{
		name: "jinny.jinny",
		email: "jinny.jinny@porori.com",
	}

	sendNotification(&jinny)
}

func sendNotification(n notifier)  {
	n.notify()
}
