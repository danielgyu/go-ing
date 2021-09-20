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

func changeAfterReturn(x int) (y int) {
	defer func() {
		y = 1111
	}()

	y += x
	return
}

func main() {
	useDefer()
	fmt.Println("")
	foo := useDeferWithReturn()
	fmt.Printf("%s from useDeferWithReturn", foo)
	fmt.Println("")
	//divideZero()
	res := changeAfterReturn(3)
	fmt.Println("res: ", res)
}
