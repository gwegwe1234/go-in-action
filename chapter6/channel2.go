package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	// 두명의 선수
	go player("Sunny", court)
	go player("Cindy", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 공이 되돌아 올 때까지 기다린다.
		ball, ok := <-court
		if !ok {
			// 채널이 닫혔으면 승리한 것으로 간주한다.
			fmt.Printf("%s 선수가 승리했습니다! \n", name)
			return
		}

		// 랜덤 값을 이용해 공을 받아치지 못했는지 확인
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 받아치지 못했습니다.\n", name)

			// 채널을 닫아 현재 선수가 패배했음을 알린다.
			close(court)
			return
		}

		// 선수가 공을 받아친 횟수를 출력하고 그 값을 증가시킨다.
		fmt.Printf("%s 선수가 %d 번째 공을 받아 쳤습니다. \n", name, ball)
		ball++

		// 공을 상대에게 보낸다.
		court <- ball
	}
}
