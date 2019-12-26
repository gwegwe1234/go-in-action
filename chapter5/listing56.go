// 인터페이스와 임베딩의 관계를 설명하기 위한 예제
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

type admin struct {
	user
	level string
}

func main()  {
	// admin 타입의 사용자를 생성
	bennie := admin {
		user : user{
			name : "bennie.maru",
			email : "bennie.maru@ddeokboggi.com",

		},
		level : "super",
	}

	// 이 경우 포함된 내부 타입이 구현한 인터페이스가 외부 타입으로 승격된다.
	sendNotification(&bennie)
}

func sendNotification(n notifier) {
	n.notify()
}
