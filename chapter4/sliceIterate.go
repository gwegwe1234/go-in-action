package main

import "fmt"

func main()  {
	slice := []int{10, 20, 30, 40}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	for i, value := range slice {
		fmt.Printf("값 : %d 값의 주소 : %x 원소의 주소 : %x\n", value, &value, &slice[i])
	}
}