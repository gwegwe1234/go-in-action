package main

import "fmt"

func main()  {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 세 번째 원소를 잘라낸다. 이때 용량을 설정한다. 길이 및 용량이 1로 지정된다.
	slice := source[2:3:3]

	fmt.Println(slice)

	// 슬라이스에 새로운 문자열을 추가한다
	slice = append(slice, "Kiwi")

	fmt.Println(slice)
	fmt.Println(source)
}
