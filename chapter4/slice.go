package main

import "fmt"

func main()  {
	var slice []int

	slice2 := make([]int, 0)
	slice = append(slice2, 0, 3, 4)
	fmt.Println(slice)
	fmt.Println(append(slice2, 0, 3, 4))
}
