package functions

import (
	"fmt"
	"time"
)

// CountDown decrase count 1 second over 0
// Blocking Timer
func CountDown(sec int) {
	for sec > 0 {
		fmt.Println(sec)
		time.Sleep(time.Second)
		sec--
	}
}
