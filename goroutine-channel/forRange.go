package main

import (
	"fmt"
	"time"
)

func main() {
	retrieve(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func retrieve(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
