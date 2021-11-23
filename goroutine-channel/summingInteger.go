package main

import "fmt"

func main() {
	ch := make(chan int)
	go sum(10, 15, ch)
	fmt.Println(<-ch)
}

func sum(x, y int, ch chan int) {
	ch <- x + y
}
