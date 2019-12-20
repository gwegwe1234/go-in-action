package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string)  {
	// 검색할 피드의 목록 조회
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 채널을 생성하여 화면에 표시할 검색 결과를 전달 받는다
	results := make(chan *Result)

	// 모든 피드를 처리할 때까지 기다릴 대기 그룹(Wait group)을 설정
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 함
	// 고루틴의 개수를 설정
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행
	for _, feed :=  range feeds {
		// 검색을 위해 검색기 조회
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴을 실행
	go func() {
		// 모든 자업이 처리될 때까지 기다린다
		waitGroup.Wait()

		// Display 함수에게 프로그램을 종류할 수 있음을 알리기 위해 채널을 닫는다.
		close(results)
	}()

	// 검색 결과 화면 표시
	Display(results)
}

// 프로그램에서 사용할 검색기를 등롣할 함수를 정의한다
func Register(feedType string, matcher Matcher)  {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다")
	}

	log.Println("등록 완료 : ", feedType, " 검색기")
	matchers[feedType] = matcher
}