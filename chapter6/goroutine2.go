package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Start Goroutine!")

	go printPrime("A")
	go printPrime("B")

	wg.Wait()
	fmt.Println("Done")
}

func printPrime(prefix string)  {
	defer wg.Done()

	next:
		for outer := 2; outer < 100000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}

		fmt.Println("완료: ", prefix)

}
