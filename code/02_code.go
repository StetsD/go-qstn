package code

import (
	"fmt"
)

// Какая проблема в коде ?
// Как можно решить ?
func Code_02() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	fmt.Println(counter)
}

// sync.WaitGroup
