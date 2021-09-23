package main

import "fmt"

type Vehicle struct {
	age int
}

type Sedan struct {
	name  string
	price float64
	int
	Vehicle
}

func main() {
	s := new(Sedan)
	fmt.Println(s)

	s.name = "my sedan"
	s.price = 600.50
	s.int = 55
	s.age = 90
	fmt.Println(s)
}
