package functions

import "fmt"

func literal() {
	printHello := func() {
		fmt.Println("hello!")
	}
	printHello()
}
