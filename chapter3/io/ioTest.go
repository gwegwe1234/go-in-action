package main

import (
	"fmt"
	"github.com/webgenie/go-in-action/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("파일을 열 때 오류 발생", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("%d개의 다넝 발견", count)
}
