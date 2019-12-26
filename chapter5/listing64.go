package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("기간 : %d", *d)
}

func main()  {
	duration(42).pretty()

	//src/github.com/gwegwe1234/go-in-action/chapter5/listing64.go:12:14: cannot call pointer method on duration(42)
	//src/github.com/gwegwe1234/go-in-action/chapter5/listing64.go:12:14: cannot take the address of duration(42)
}
