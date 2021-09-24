package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    []string
	price    []float64
	quantity []int
}

func main() {
	file1, _ := os.Open("csvSample.csv")
	defer file1.Close()

	useReader(file1)
	fmt.Println()

	file2, _ := os.Open("csvSample.csv")
	defer file2.Close()
	useScanner(file2)
}

func useReader(file *os.File) {
	reader := bufio.NewReader(file)
	book := new(Book)
	for {
		read, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		split := strings.Split(read, ",")
		splitAndAppend(split, book)
	}
	printBook(book)
}

func useScanner(file *os.File) {
	s := bufio.NewScanner(file)
	book := new(Book)

	for s.Scan() {
		read_line := s.Text()
		split := strings.Split(read_line, ",")
		splitAndAppend(split, book)
	}
	printBook(book)
}

func splitAndAppend(split []string, book *Book) {
	price, _ := strconv.ParseFloat(split[1], 32)
	quantity, err := strconv.Atoi(split[2])
	if err != nil {
		fmt.Println("err:", err)
	}

	book.title = append(book.title, split[0])
	book.price = append(book.price, price)
	book.quantity = append(book.quantity, quantity)
}

func printBook(book *Book) {
	fmt.Println("title:", book.title)
	fmt.Println("price:", book.price)
	fmt.Println("quantity:", book.quantity)
}
