package gorutinelessons

import (
	"fmt"
	"sync"
	"time"
)

func MainerGo6() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			work()
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func work() {
	fmt.Println("Goroutine start")
	time.Sleep(time.Second * 2)
	fmt.Println("Goroutine stop")

}
