package gorutinelessons

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func MainerGo210() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goTest(i, wg)
	}
	wg.Wait()
}

func goTest(id int, wg *sync.WaitGroup) {
	for i := 0; i < 2; i++ {
		second := strings.Split(
			time.Now().Format("Jan _2 15:04:05.000"), " ")[2]
		time.Sleep(time.Second)
		fmt.Printf("Поток %d выполнился со счётчиком %d, во время %v\n",
			id, i, second)
	}
	wg.Done()
}
