package main

import "fmt"

func main() {
	var array [5]int
	array2 := [5]int{1, 2, 3, 4, 5}
	array3 := [...]int{1, 2, 3, 4, 5}
	array4 := [5]int{1: 1, 2: 2}

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	array5 := [5]int{10, 20, 30, 40, 50}

	array5[2] = 25

	fmt.Println(array5)

	array6 := [5]*int{0: new(int), 1: new(int)}
	*array6[0] = 10
	*array6[1] = 20

	array7 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	var array8 [5]string
	array8 = array7

	fmt.Println(array8)

	var array9 [3]*string

	array10 := [3]*string{new(string), new(string), new(string)}

	*array10[0] = "Red"
	*array10[1] = "Blue"
	*array10[2] = "Green"

	array9 = array10

	fmt.Println(*array9[0])
}
