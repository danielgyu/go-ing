package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return n * factorial(n-1)
}

func main() {
	res := factorial(10)
	fmt.Println(res)
}
