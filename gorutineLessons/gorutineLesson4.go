package gorutinelessons

import "fmt"

func myFunc() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		close(done)
	}()
	return done
}

func MainerGo4() {
	<-myFunc()
}
