package main

import (
	_ "github.com/gwegwe1234/go-in-action/chapter2/sample/matchers"
	"github.com/gwegwe1234/go-in-action/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	// 표준 출력으로 로그를 출력하도록 변경한다
	log.SetOutput(os.Stdout)
}

func main() {
	// 지정된 검색어로 검색을 수행
	search.Run("Sherlock Holmes")
}
