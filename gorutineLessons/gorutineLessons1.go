package gorutinelessons

import "fmt"

func task(c chan int, N int) {
	c <- N + 1
}

func MainerGo() {
	channel := make(chan int, 1)
	task(channel, 3)
	fmt.Println(<-channel)
}
