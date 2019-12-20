package search

import (
	"encoding/json"
	"os"
)

const dataFile = "./src/github.com/gwegwe1234/go-in-action/chapter2/sample/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error)  {
	// 파일을 연다
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer 함수를 이용해 이 함수가 리턴딜 때 앞서 열어둔 파일이 꼭 닫히도록 해준다
	defer file.Close()

	// 파일을 읽어 Feed 구조체의 포인터의 슬라이스로 변환
	var feeds []*Feed
	//feeds2 := new([]*Feed)
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
