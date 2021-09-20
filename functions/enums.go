package main

import "fmt"

type Gender int

const (
	UNKNOWN = iota
	FEMALE
	MALE
)

func main() {
	fmt.Println(UNKNOWN)
	fmt.Println(MALE)
	fmt.Println(FEMALE)
}
