package gorutinelessons

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MainerGo9() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 5
	for i := 0; i < n; i++ {
		in1 <- i
		in2 <- 10 * (i + 1)
	}
	merge2Channels(fn, in1, in2, out, n)
	fmt.Println("not blocked if printed first")
	time.Sleep(time.Second * 3)
	for i := 0; i < n; i++ {
		fmt.Println(<-out)
	}
}

func fn(fx int) int {
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Microsecond)
	return fx
}

func merge2Channels(fn func(int) int,
	in1 <-chan int,
	in2 <-chan int,
	out chan<- int,
	n int) {
	first_channel, second_channel, mu, wg := make(chan int), make(chan int), new(sync.Mutex), new(sync.WaitGroup)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			go func() {
				fx, fx2 := <-in1, <-in2
				first_channel <- fx
				second_channel <- fx2
			}()
			mu.Unlock()
		}(wg, mu)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			wg.Wait()
			mu.Lock()
			out <- <-first_channel + <-second_channel
			mu.Unlock()
		}(wg, mu)
	}
}
