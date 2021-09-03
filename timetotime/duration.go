package timetotime

import (
	"fmt"
	"strings"
	"time"
)

func Duration() {
	// for Stepic
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	result := a + " " + b + " " + c
	// fmt.Println(result)
	//
	// Для отладки
	// result := "13.03.2018 14:00:15,12.03.2018 14:00:15"
	//
	dateOne, dateTwo := strings.Split(result, ",")[0], strings.Split(result, ",")[1]
	timeDateOne, _ := time.Parse("02.01.2006 15:04:05", dateOne)
	timeDateTwo, _ := time.Parse("02.01.2006 15:04:05", dateTwo)
	if timeDateOne.Format("02.01.2006 15:04:05") >
		timeDateTwo.Format("02.01.2006 15:04:05") {
		fmt.Println(time.Since(timeDateTwo).Round(time.Second) -
			time.Since(timeDateOne).Round(time.Second))
	} else {
		fmt.Println(time.Since(timeDateOne).Round(time.Second) -
			time.Since(timeDateTwo).Round(time.Second))
	}
	fmt.Println("8 program")
}
