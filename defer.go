package main

import "fmt"

func useDefer() {
	defer fmt.Println("useDefer ended")
	fmt.Println("deferring function")
}

func useDeferWithReturn() string {
	defer fmt.Println("defer with return")
	fmt.Println("deferring with return")
	return "foo"
}

func divideZero() {
	defer fmt.Println("deferring divide zero")
	zero := 0
	fmt.Println(1 / zero)
}

func main() {
	useDefer()
	fmt.Println("")
	foo := useDeferWithReturn()
	fmt.Printf("%s from useDeferWithReturn", foo)
	fmt.Println("")
	divideZero()
}
