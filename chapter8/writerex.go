package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Buffer 값을 생성한 후 버퍼에 문자열을 출력
	// 이때 io.Writer 인터페이스를 구현한 Write 메소드 호출
	var b bytes.Buffer
	b.Write([]byte("할로"))

	// 버퍼에 문자여을 결합하기 우해 Fprintf 함수를 호출
	// 이때 bytes.Buffer의 주소를 io.Writer 타입 매개변수로 전달.
	fmt.Fprintf(&b, "고랭!")

	// 버퍼의 표준 출력 장치에 쓴다.
	b.WriteTo(os.Stdout)
}
