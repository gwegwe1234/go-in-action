package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 모든 고루틴이 값을 증가라혀고 시도하는 변수
	counter int

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("최종결과 : ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := count

		// 스레드를 양보하여 큐로 돌아가도록 한다.
		runtime.Gosched()

		value++

		counter = value
	}
}
