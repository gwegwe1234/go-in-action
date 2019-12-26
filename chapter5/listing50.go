// 타입 임베딩을 이용해 다른 타입을 포함하는 방법과
// 이 경우 내부 및 외부 타입의 관계를 확인하기 위한 예제 프로그램
package main

import "fmt"

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

	// 내부 타입의 메소드를 바로 호출할 수 있다.
	bennie.user.notify()

	// 내부 타입의 메소드가 승격 되었다.
	bennie.notify()
}
