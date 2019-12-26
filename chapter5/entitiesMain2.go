package main

import (
	"fmt"
	"github.com/webgenie/go-in-action/chapter5/listing74/entities"
)

func main()  {
	a := entities.Admin{
		Rights: 3,
	}

	// 비노출 타입인 내부 타입의 노출 타입 필드를 세팅
	a.Name = "admin"
	a.Email = "Email"

	fmt.Println(a.Name)
	fmt.Println(a.Email)
	fmt.Println(a.Rights)
}
