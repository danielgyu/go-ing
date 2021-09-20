package main

import "fmt"

func closure() func(x, y int) int {
	fmt.Println("Wrapping function")
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	fmt.Println("vim-go")
	adder := closure()
	fmt.Println(adder(1, 2))
}
