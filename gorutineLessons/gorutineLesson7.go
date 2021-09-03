package gorutinelessons

import "fmt"

func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	chant := make(chan int, 3)

	go func() {
		defer close(chant)
		select {
		case f1 := <-firstChan:
			fm := f1 * f1
			chant <- fm
		case f2 := <-secondChan:
			fm := f2 * 3
			chant <- fm
		case <-stopChan:
			return
		}
	}()
	return chant
}

func MainerGo7() {
	ch1, ch2 := make(chan int, 3), make(chan int, 3)
	stop := make(chan struct{})
	ch1 <- 15
	r := calculator1(ch1, ch2, stop)
	fmt.Println(<-r)
}
