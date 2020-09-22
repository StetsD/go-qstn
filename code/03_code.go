package code

import (
	"fmt"
	"time"
)

func Code_03() {
	words := []string{"foo", "bar", "baz"}
	done := make(chan bool)
	defer close(done)
	for _, word := range words {
		go func(word string) {
			time.Sleep(1 * time.Second)
			fmt.Println(word)
			done <- true
		}(word)
	}

	for range words {
		<-done
	}
}
