package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	useBufioReader()
	useIoReader()
	useIoutilsReader()
	useBufferedReader()
}

func useBufioReader() {
	file, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		lineString, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("err")
			return
		}
		fmt.Println("from BufReader:", lineString)
	}
}

func useIoReader() {
	//file, err := os.Open("sample.txt", os.O_RDONLY, 0)
	file, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return
	}
	fmt.Println("from ioReader:", string(contents))
}

func useIoutilsReader() {
	buf, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		return
	}
	fmt.Println("from ioutil:", string(buf))
}

func useBufferedReader() {
	buf := make([]byte, 1024)
	fileHandler, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(fileHandler)

	for {
		read, err := reader.Read(buf)
		if err != nil {
			return
		}
		if read == 0 {
			break
		}
		fmt.Println("from bufferedReader:", string(buf))
	}
}
