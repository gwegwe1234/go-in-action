package main

import (
	"github.com/gwegwe1234/go-in-action/chapter7/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGroutines    = 25
	pooledResources = 2
)

// 공유 자원을 표현한 구조체
type dbConnection struct {
	ID int32
}

// dbConnection이 Pool에 의해 관리될 수 있도록 io.Closer 인터페이스 구현
func (dbConn *dbConnection) Close() error {
	log.Println("닫힘 : 데이터베이스 연결", dbConn.ID)
	return nil
}

// 각 데이터베이스에 유일한 id를 할당하기 위한 변수
var idCounter int32

// 풀이 새로운 리소스가 필요할 때 호출할 팩토리 메소드
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)

	log.Println("생성 : 새 데이터베이스 연결", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGroutines)

	// 데이터베이스 연결을 관리할 풀을 생성
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 풀에서 데이터베이스 연결을 가져와 쿼리 실행
	for query := 0; query < maxGroutines; query++ {
		// 각 고루틴에는 질의 값의 복사본을 전달
		// 그렇지 않으면 고루틴들이 동일한 쿼리값을 공유
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	// 풀을 닫는다.
	log.Println("프로그램 종료")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	// 풀에서 데이터베이스 연결 리소스 획득
	conn, err := p.Acquire()

	if err != nil {
		log.Println(err)
		return
	}

	// 데이터베이스 연결 리소스를 다시 풀로 되돌린다.
	defer p.Release(conn)

	// 질의문이 실행되는 것처럼 얼마 동안 대기한다.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("질의: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
