package main

import "log"

func init() {
	log.SetPrefix("추적 : ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 함수는 표준 로거에 메세지 출력
	log.Println("메세지")

	// Fatalln 함수는 Println() 함수를 실행한 후 os.Exit(1)을 추가 호출
	log.Fatalln("아주 치명적인 오류메세지!")

	// Panicln함수는 Println 함수를 호출한 후 panic 함수를 추가 호출
	log.Panicln("패닉 메세지!")
}
