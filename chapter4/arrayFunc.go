package main

import "fmt"

func main()  {
	var array [1e6]int
	foo(array)

}

func foo(array [1e6]int)  {
	fmt.Println(array)
}
