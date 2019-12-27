package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	nuberGoroutines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().Unix())
}

func main()  {
	// 버퍼가 있는 채널 생성
	tasks := make(chan string, taskLoad)

	wg.Add(nuberGoroutines)

	// 작업을 처리할 고루틴을 실행한다.
	for gr := 1 ; gr <= nuberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 실행할 작업을 추가한다.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업 : %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 작업이 할당될 때까지 대기한다.
		task, ok := <- tasks

		if !ok {
			// 채널이 닫힌 경우
			fmt.Printf("작업자 : %d : 종료합니다. \n", worker)
			return
		}

		fmt.Printf("작업자 : %d : 작업시작 : %s\n", worker, task)

		sleep := rand.Intn(100)

		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("작업자 : %d : 작업 완료 : %s\n", worker, task)
	}
}
