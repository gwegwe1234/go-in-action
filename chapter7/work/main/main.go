package main

import (
	"github.com/gwegwe1234/go-in-action/chapter7/work"
	"log"
	"sync"
	"time"
)

// 화면에 출력할 이름들을 슬라이스로 선언
var names = []string{
	"ted",
	"sunny",
	"cindy",
	"jade",
	"liam",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 이름 슬라이스르 반복
		for _, name := range names {
			// namePrinter 변수를 선언하고 이름을 지정
			np := namePrinter{
				name: name,
			}

			go func() {
				// 실행할 작업 등록
				// Run 메소드가 리턴되면 해당 작업이 처리된 걸로 간주
				p.Run(&np)

				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 작업 풀을 종료하고 이미 등록된 작업들이 종료될 때까지 대기
	p.Shutdown()
}
