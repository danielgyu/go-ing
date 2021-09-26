package main

import (
	"fmt"
)

func main() {
	stringChan := make(chan string)
	go sendData(stringChan)
	//go receiveData(stringChan)
	//time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	ch <- "from"
	ch <- "Lee"
	fmt.Println("@@@finished sending")
}

func receiveData(ch chan string) {
	defer close(ch)
	var input string
	for {
		input = <-ch
		fmt.Println("received data", input)
	}
}
