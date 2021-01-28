package goroutines

import "fmt"

func simpleChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for n := range c {
		fmt.Println(n)
	}
}
