package main

import "fmt"

type Person struct {
	name string
}

func (p Person) String() string {
	return "hello i am a person"
}

func main() {
	fmt.Println(Person{"lee"})
}
