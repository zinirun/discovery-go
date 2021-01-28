package goroutines

import "fmt"

func plusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func exPlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range plusOne(plusOne(c)) {
		fmt.Println(num)
	}
}

// IntPipe type
type IntPipe func(<-chan int) <-chan int

// Chain makes pipeline goroutines
func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}
