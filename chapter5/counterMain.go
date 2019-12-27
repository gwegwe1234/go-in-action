// 다음 패키지에서 비노출 식별자에 대한 접근이 차단되는것을 보여주기 위한 프로그램
package main

import (
	"fmt"
	"github.com/gwegwe1234/go-in-action/chapter5/counters"
)

func main() {
	//counter := counters.alertCounter(10)

	counter := counters.New(10)
	fmt.Println(counter)
	//src/github.com/gwegwe1234/go-in-action/chapter5/counterMain.go:10:13: cannot refer to unexported name counters.alertCounter
	//src/github.com/gwegwe1234/go-in-action/chapter5/counterMain.go:10:13: undefined: counters.alertCounter
}
