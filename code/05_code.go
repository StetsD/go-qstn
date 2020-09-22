package code

import (
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		wg.Done()
		ch <- 1
	}()

	return ch
}

func Code_05() {
	timeStart := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			<-worker(&wg)
		}()
	}

	wg.Wait()
	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
}
