package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	slice[1] = 15

	fmt.Println(slice)

	newSlice := slice[1:3]

	fmt.Println(newSlice)

	newSlice[1] = 35

	fmt.Println(slice)
	fmt.Println(newSlice)
}
