package publicer

import "fmt"

func NewArray() []int {
	array := make([]int, 10, 12)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Println("7 program")
	return array
}
