package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (u *admin) notify() {
	fmt.Printf("관리자에게 메일을 전송합니다: %s<%s>\n", u.name, u.email)
}

func main() {
	// admin 타입의 사용자를 생성
	bennie := admin{
		user: user{
			name:  "bennie.maru",
			email: "bennie.maru@ddeokboggi.com",
		},
		level: "super",
	}

	// 이 경우 포함된 내부 타입이 구현한 인터페이스가 외부 타입으로 승격되지 않는다
	sendNotification(&bennie)

	// 내부 메소드에 직접 접근할 수 있다.
	bennie.user.notify()

	// 이 경우 내부 타입의 메소드가 승격되지 않는다
	bennie.notify()
}

func sendNotification(n notifier) {
	n.notify()
}
