package interfaces

// 인터페이스는 일반적으로 ~er로 명명
type methoder interface {
	Method1()
	Method2(i int) error
}

type reader interface {
	Method3()
	Method4(s string) error
}

// 인터페이스끼리 합칠 수 있음
type integrater interface {
	methoder
	reader
}
