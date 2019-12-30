package main

import (
	"github.com/gwegwe1234/go-in-action/chapter7/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("작업을 시작합니다")

	// 실행 시간을 이용해 새로운 작업 실행기 생성
	r := runner.New(timeout)

	// 수행할 작업 등록
	r.Add(createTask(), createTask(), createTask())

	// 작업을 실행하고 결과를 처리
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("지정된 작업 시간 초과")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("운영체제 인터럽트 발생")
			os.Exit(2)
		}
	}

	log.Println("프로그램 종료")
}

// 지정된 시간동안 암것도 안하는 예제를 위한 생성함수
func createTask() func(int) {
	return func(id int) {
		log.Printf("프로세서 - 작업 #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
