package goroutines

import (
	"fmt"
	"time"
)

//FanOut Example
func FanOut() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	defer close(c)
	/*
		채널을 굳이 닫을 필요는 없었지만 고루틴이 종료되지 않으면 메모리 누수가 발생
		또한 채널을 닫는 것은 방송(broadcast) 효과가 있음 -> 신호를 모두에게 전달하기 위한 깔끔한 방법
	*/
}
