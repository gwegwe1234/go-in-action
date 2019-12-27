package main

import (
	"fmt"
	"github.com/gwegwe1234/go-in-action/chapter5/entities"
)

func main() {
	// entities 패키지내의 User 구조체 가져온다.

	riley := entities.User{
		Name:  "Riley.life",
		email: "Riley.life@email.com",
	}

	fmt.Println(riley)

	// entitiesMain.go:13:3: unknown field 'email' in struct literal of type entities.User
}
