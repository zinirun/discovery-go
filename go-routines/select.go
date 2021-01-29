package goroutines

import (
	"fmt"
	"time"
)

func selectExample(c chan rune) {
	// 채널에 값이 있으면 받고, 없으면 스킵하는 흐름
	select {
	case n := <-c:
		fmt.Println(n)
	default:
		fmt.Println("Data is not ready. Skipping..")
	}
}

func selectTimerExample(recv chan rune) {
	send := make(chan rune)
	// 채널과의 통신을 기다리되 일정 시간동안만 대기 -> time.After 사용
	select {
	case n := <-recv:
		fmt.Println(n)
	case send <- 1:
		fmt.Println("sent 1")
	case <-time.After(5 * time.Second):
		fmt.Println("No data for 5 seconds")
		return
	}
}
