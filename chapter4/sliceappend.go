package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	newSlice = append(newSlice, 60)

	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice2 := append(slice, 60)
	fmt.Println(newSlice2)
}
