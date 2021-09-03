package interfaceslesson

import (
	"fmt"
	"os"
)

func Interfaces_lesson() {
	var (
		value1    interface{} = 14.1
		value2    interface{} = 12.3
		operation interface{} = "+"
	)
	typeValue1 := fmt.Sprintf("%T", value1)
	typeValue2 := fmt.Sprintf("%T", value2)
	var1, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v:%T", value1, value1)
		os.Exit(0)
	}
	var2, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v:%T", value2, value2)
		os.Exit(0)
	}
	var result float64
	if typeValue1 == "float64" && typeValue2 == "float64" {
		switch operation {
		case "-":
			result = float64(var1) - float64(var2)
			fmt.Printf("%.4f", float64(result))
		case "+":
			result = float64(var1) + float64(var2)
			fmt.Printf("%.4f", float64(result))
		case "*":
			result = float64(var1) * float64(var2)
			fmt.Printf("%.4f", float64(result))
		case "/":
			result = float64(var1) / float64(var2)
			fmt.Printf("%.4f", float64(result))
		default:
			fmt.Println("неизвестная операция")
			os.Exit(0)
		}
		os.Exit(0)
	}
	fmt.Println("2 program")
}
