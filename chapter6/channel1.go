package main

import "fmt"

func main() {
	// 버퍼의 크기가 정해지지 않은 정수 채널
	unbuffered := make(chan int)

	// 버퍼의 크기가 정해진 문자열 채널
	buffered := make(chan string, 10)

	buffered <- "Gopher"

	value := <-buffered

	fmt.Printf(value)
}
