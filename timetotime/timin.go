package timetotime

import (
	"fmt"
	"time"
)

func TimEr() {
	var a string = "1986-04-16T05:20:00+06:00"
	fmt.Println(a)
	fileTimer, _ := time.Parse(time.RFC3339, a)
	fmt.Println(fileTimer.Format("Mon Jan _2 15:04:05 MST 2006"))
	fmt.Println("10 program")
}
