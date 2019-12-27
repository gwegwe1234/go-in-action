package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main()  {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	// 경기 시작!
	baton <- 1

	wg.Wait()
}

func Runner(baton chan int)  {
	var newRunner int

	// 바통을 전달 받을 때 까지 기다린다.
	runner := <- baton

	// 트랙을 달린다.
	fmt.Printf("%d 번째 주자가 바통을 받아 달리기 시작합니다! \n", runner)

	// 새로운 주자가 교체 지점에서 대기
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("%d 번째 주자가 대기합니다. \n", newRunner)
		go Runner(baton)
	}

	// 트랙을 달린다
	time.Sleep(1000 * time.Millisecond)

	// 경기가 끝났는지 확인
	if runner == 4 {
		fmt.Printf("%d 번째 주자가 도착했습니다. 경기가 끝났습니다.", runner)
		wg.Done()
		return
	}

	// 다음 주자에게 바통을 넘긴다.
	fmt.Printf("%d 번째 주자가 %d 번째 주자에게 바통을 넘겨줍니다.\n", runner, newRunner)

	baton <- newRunner
}