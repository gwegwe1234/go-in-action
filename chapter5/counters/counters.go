package counters

// 알림 횟수를 저장하기 위한 정수 값을 저장하는 alertCounter 타입을 비노출 타입으로 선언
type alertCounter int

// New라는 함수로 alertCounter 타입을 리턴한다.
func New(value int) alertCounter {
	return alertCounter(value)
}
