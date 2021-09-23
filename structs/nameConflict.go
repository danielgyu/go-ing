package main

import "fmt"

type Vehicle struct {
	name string
}

type Sedan struct {
	name int
	Vehicle
}

func main() {
	s := Sedan{10, Vehicle{"vehicle name"}}
	fmt.Println(s.name)
}
