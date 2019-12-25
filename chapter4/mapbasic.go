package main

import "fmt"

func main()  {
	// string key , int value
	dict := make(map[string]int)

	//string key, string value
	dict2 := map[string]string{"Red" : "red", "Orange" : "orange"}

	// 빈 맵 생성
	colors := map[string]string{}

	// 키, 밸류 설정해 값넣기
	colors["Red"] = "red"

	value, exists := colors["Blue"]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key is not exist")
	}

	// 맵을 선언하기만 하면 nil 맵이 설정된다
	var color map[string]string
}
