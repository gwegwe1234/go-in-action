package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 구조체는 여러 개의 고루틴에서 안전하게 공유하기 위한 리소스 집합 관리
// 이 풀에서 관리하기 위한 리소스는 io.Closer 인터페이스를 반드시 구현해야한다.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 에러는 리소스를 획득하려 할 때 풀이 닫혀 있는 경우에 발생
var ErrPoolClosed = errors.New("풀이 닫혔습니다.")

// New 함수는 리소스 관리 풀 생성
// 풀은 새로운 리소스를 할당하기 위한 함수와 풀의 크기를 매개변수로 정의
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("풀의 크기가 너무 작습니다.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 풀에서 리소스를 획득하는 메소드
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 사용 가능한 리소스가 있는지 확인
	case r, ok := <-p.resources:
		log.Println("리소스 획득:", "공유된 리소스")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 사용 가능한 리소스가 없으면 새로운 리소스 생성
	default:
		log.Println("리소스 획득:", "새로운 리소스")
		return p.factory()
	}
}

// 풀에 리소스를 반환하는 메소드
func (p *Pool) Release(r io.Closer) {
	// 안전한 작업을 위해 잠금 설정
	p.m.Lock()
	defer p.m.Unlock()

	// 풀이 닫혔으면 리소스 해제
	if p.closed {
		r.Close()
		return
	}

	select {
	// 새로운 리소스를 큐에 추가
	case p.resources <- r:
		log.Println("리소스 반환:", "리소스 큐에 반환")

	// 리소스 큐가 가득 찬 경우 리소스를 해제
	default:
		log.Println("리소스 반환:", "리소스 해제")
		r.Close()
	}
}

// 풀을 종료하고 생섣ㅇ된 모든 리소스를 해제하는 메소드
func (p *Pool) Close() {
	// 안전한 작업을 위해 잠금 설정
	p.m.Lock()
	defer p.m.Unlock()

	// 풀이 이미 닫혔으면 암것도 안함
	if p.closed {
		return
	}

	// 풀을 닫힌 상태로 전환
	p.closed = true

	// 리소스 해제 전에 채널을 닫는다.
	// 데드락 걸릴 가능성이 있으므로
	close(p.resources)

	// 리소스 해지
	for r := range p.resources {
		r.Close()
	}
}
