package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Scanln(&name)
	fmt.Println("name is", name)

	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter input:")

	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("Input was:", input)
	}
}
