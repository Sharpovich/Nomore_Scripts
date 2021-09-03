package gorutinelessons

import (
	"fmt"
)

type dome struct{}

func MainerGo8() {
	c := make(chan int)
	done := make(chan struct{}, 1)
	r := calculator(c, done)
	c <- 255
	c <- 1
	c <- 2
	c <- 3
	done <- dome{}

	fmt.Println(<-r)
}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		var sum int = 0
		for {
			select {
			case <-done:
				ch <- sum
				return
			case fn := <-arguments:
				sum += fn
			}
		}
	}()
	return ch
}
