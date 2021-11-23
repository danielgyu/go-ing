package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{100, 101, 102}
	intChan := make(chan int)

	/* non-blocking
	go func(ch chan int) {
		for n := range ch {
			fmt.Printf("received %d\n", n)
		}
	}(intChan)
	*/

	for _, num := range nums {
		go func(n int, ch chan int) {
			time.Sleep(time.Second * 3)
			fmt.Printf("n is %d\n", n)
			ch <- n
		}(num, intChan)
	}
	fmt.Println("sent all goroutines")

	for range nums {
		n := <-intChan
		fmt.Printf("received %d\n", n)
	}
}
