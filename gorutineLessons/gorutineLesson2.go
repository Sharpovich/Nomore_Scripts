package gorutinelessons

import "fmt"

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func MainerGo2() {
	channel := make(chan string, 5)
	task2(channel, "F")
	fmt.Println(<-channel)
}
