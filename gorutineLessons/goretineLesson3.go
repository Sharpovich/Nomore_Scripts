package gorutinelessons

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var firstValue string
	for v := range inputStream {
		if firstValue != v {
			outputStream <- v
			firstValue = v
			continue
		} else {
			continue
		}
	}
	close(outputStream)
}

func MainerGo3() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
}
