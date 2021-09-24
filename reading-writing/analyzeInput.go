package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input:")

	input, err := inputReader.ReadString('S')

	if err != nil {
		return
	}

	fmt.Println("Number of characters:", len(input))
	fmt.Println("Number of words:", len(strings.Split(input, " ")))

	buffer := make([]byte, 32*1024)
	count := 0
	lineSeparator := []byte{'\n'}

	c, err := inputReader.Read(buffer)
	count += bytes.Count(buffer[:c], lineSeparator)
	if err != nil {
		return
	}
	fmt.Println("Number of lines:", count)
}
