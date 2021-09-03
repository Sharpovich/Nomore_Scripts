package timetotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func TimeinTwo() {
	var dst string = "2020-05-15"
	var tst string = "08:00:00"
	// var dst, tst string
	// fmt.Scan(&dst)
	// fmt.Scan(&tst)
	result := dst + " " + tst
	fileTimer, _ := time.Parse("2006-01-02 15:04:05", result)
	hour, _ := strconv.Atoi(strings.Split(tst, ":")[0])
	minute, _ := strconv.Atoi(strings.Split(tst, ":")[1])
	current := time.Date(fileTimer.Year(),
		fileTimer.Month(),
		fileTimer.Day(),
		fileTimer.Hour(),
		fileTimer.Minute(),
		fileTimer.Second(), 0, time.Local)
	if hour >= 13 && minute >= 00 {
		current = time.Date(fileTimer.Year(),
			fileTimer.Month(),
			fileTimer.Day()+1,
			fileTimer.Hour(),
			fileTimer.Minute(),
			fileTimer.Second(), 0, time.Local)
		fmt.Println(current.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(current.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("9 program")
}
