package main

import "fmt"

type Persons []Person

type Person struct {
	firstName string
	lastName  string
}

type SortInterface interface {
	sort() Persons
}

func main() {
	fmt.Println("vim-go")
}
