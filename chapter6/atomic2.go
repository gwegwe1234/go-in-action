package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// 고루틴이 작업할 시간을 1초 줌
	time.Sleep(1 * time.Second)

	// 종료 플래그 설정
	fmt.Println("프로그램 종료!")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("작업 진행중 : %s\n", name)
		time.Sleep(250 * time.Millisecond)

		// 종료 플래그 확인
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("작업을 종료합니다: %s\n", name)
			break
		}
	}
}
