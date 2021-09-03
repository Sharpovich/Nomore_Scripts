package timetotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func UnixDateConvert() {
	var a, b, c, d string
	fmt.Scan(&a, &b, &c, &d)
	dateTime := a + " " + b + " " + c + " " + d
	const now = 1589570165
	value1, _ := strconv.Atoi(strings.Split(
		strings.Split(
			dateTime, ".")[0], " ")[0])
	value2, _ := strconv.Atoi(strings.Split(
		strings.Split(
			dateTime, ".")[1], " ")[1])
	resultSecond := time.Unix(int64(
		((value1*60)+value2)+now), 0).UTC()
	fmt.Println(resultSecond.Format("Mon Jan _2 15:04:05 MST 2006"))
	fmt.Println("11 program")
}
