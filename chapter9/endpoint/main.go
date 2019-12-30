// 간단한 웹서비스 예제
package main

import (
	"github.com/gwegwe1234/go-in-action/chapter9/endpoint/handler"
	"log"
	"net/http"
)

// 애플리케이션 진입점
func main() {
	handlers.Routes()

	log.Println("웹서비스 실행 중: 포트: 4000")
	http.ListenAndServe(":4000", nil)
}
