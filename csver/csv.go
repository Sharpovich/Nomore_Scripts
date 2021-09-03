package csver

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func CsvOuter() {
	file, _ := os.Open("csver/data.csv")
	reader := csv.NewReader(file)
	res, _ := reader.ReadAll()
	for _, v := range res {
		for _, v := range v {
			resArray := strings.Split(v, ";")
			for i, v := range resArray {
				if v == "0" {
					fmt.Println(i + 1)
				}
			}
		}
	}
	fmt.Println("1 program")
}
