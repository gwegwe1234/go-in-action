package entities

// 외부에 노출된 User 구조체 생성
type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Rights int
}
