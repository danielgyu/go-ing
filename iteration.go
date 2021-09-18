package main

import "fmt"

func main() {
	var i int = 10
	for i > 5 {
		fmt.Println("iterating...")
		i--
	}

	var str string = "hello"
	for idx, val := range str {
		fmt.Printf("idx: %d, val: %c\n", idx, val)
	}
}
