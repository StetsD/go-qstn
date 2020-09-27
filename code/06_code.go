package code

import "fmt"

func Code_06() {
	var a chan int

	a <- 23

	fmt.Println(<-a)
}
