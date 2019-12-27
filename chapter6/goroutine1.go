package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 스케쥴러가 하나의 논리 프로세서를 할당
	runtime.GOMAXPROCS(runtime.NumCPU())

	// wg는 프로그램의 종료를 대기하기 위해 사용
	// 각각의 고루틴마다 하나씩, 총 두 개의 카운트를 추가한다.

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴을 실행합니다.")

	// 익명 함수를 선언하고 고루틴을 생성한다.
	go func() {
		// main 함수에게 종료를 알리기 위한 Done 함수 호출을 예약
		defer wg.Done()

		// 알파벳을 세번 출력
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println()

	// 익명 함수를 선언하고 고루틴을 생성한다.
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 고루틴이 종료 될 때까지 대기한다.
	wg.Wait()
	fmt.Println("대기중!!")
	fmt.Println("프로그램 종료!")
}
