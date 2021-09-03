package newbattery

import "fmt"

type Stringer interface {
	String() string
}

type Battery struct {
	Charge string
}

func (b Battery) String() string {
	var count_zero, count_one int
	for i := 0; i < 10; i++ {
		if string(b.Charge[i]) == "1" {
			count_one++
		}
		if string(b.Charge[i]) == "0" {
			count_zero++
		}
	}
	var array [12]string = [12]string{"[", "", "", "", "", "", "", "", "", "", "", "]"}
	for i := 1; i < count_zero+1; i++ {
		array[i] = " "
	}
	for i := count_zero + 1; i < 12-1; i++ {
		array[i] = "X"
	}
	for i := 0; i < 12; i++ {
		fmt.Print(array[i])
	}
	fmt.Print("\n")
	var s string
	for i := 0; i < 12; i++ {

	}
	fmt.Println("5 program")
	return s
}
