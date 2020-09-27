package code

import (
	"fmt"
	//"sync"
)

// Что можно сделать, чтобы сохранить последовательность слайса с учетом параллельного перебора ?

func Code_03() {
	//mu := sync.Mutex{}
	var result []string
	words := []string{"foo", "bar", "baz"}
	done := make(chan bool)
	defer close(done)
	for _, word := range words {
		//mu.Lock()
		go func(word string) {
			result = append(result, word)
			//mu.Unlock()
			done <- true
		}(word)
	}

	for range words {
		<-done
	}

	fmt.Println(result)
}
