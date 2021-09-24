package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.OpenFile("writeSample.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	content := "Hello World!\n"

	writer.WriteString(content)
	writer.Flush()
}
