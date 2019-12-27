package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup

	mutex sync.Mutex
)

func main() {

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Result : %d", counter)

}

func incCounter(id int) {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		//이 임계 지역에는 한번에 하나의 고루틴만 접근 가능
		mutex.Lock()
		{
			value := counter

			runtime.Gosched()

			value++

			counter = value
		}
		mutex.Unlock()
	}
}
