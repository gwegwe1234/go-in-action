package main

import "fmt"

// 프로그램의 사용자를 표현하는 user 타입
type user struct {
	name string
	email string
}

// 값 수신자와 함께 notify 메소드를 선언한다.
func (u user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다 : %s<%s>\n", u.name, u.email)
}

// 포인터 수신자와 함께 changeEmail 메소드를 선언
func (u *user) changeEmail(email string) {
	u.email = email
}

func main()  {
	// user 타입의 값을 이용하여 값 수신자(receiver)에 선언한 메소드 호출
	ted := user{"ted", "ted.park@dododada.net"}
	ted.notify()

	// user 타입의 포인터를 이용하여 값 수신자에 선언한 메소드 호출
	sunny := &user{"sunny", "sunny.sunny@tototata.net"}
	sunny.notify()

	ted.changeEmail("ted.park@kakaka.com")
	ted.notify()

	sunny.changeEmail("sunny.sunny@mamamam.co.kr")
	sunny.notify()
}
