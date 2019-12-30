package basictest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T)  {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("콘텐츠 다운로드 기능 테스트 시작!")
	{
		t.Logf("\tUrl \"%s\" 호출 시 상태 코드가 \"%d\"인지 확인.", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tHTTP GET 요청을 보냈는지 확인", ballotX, err)
			}
			t.Log("\t\tHTTP GET 요청을 보냈는지 확인", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t상태코드가 \"%d\"인지 확인 %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
