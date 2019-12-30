package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("에러 로그 파일을 열 수 없습니다.", err)
	}

	Trace = log.New(ioutil.Discard,
		"추적 : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"정보 : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"경고 : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"에러 : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("트레이스 로그메세지")
	Info.Println("인포 로그메세지")
	Warning.Println("워닝 로그메세지")
	Error.Println("에러 로그메세지")
}
