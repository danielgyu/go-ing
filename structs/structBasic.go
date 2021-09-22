package main

import "fmt"

type myStruct struct {
	name string
	age  int
}

func main() {
	person1 := new(myStruct)
	fmt.Println(person1)
	fmt.Println(person1.name)
	fmt.Println(person1.age)
	fmt.Println()

	person2 := myStruct{"lee", 10}
	fmt.Println(person2)
	fmt.Println(person2.name)
	fmt.Println(person2.age)
	fmt.Println()

	person3 := &myStruct{"kim", 20}
	fmt.Println(person3)
	fmt.Println(person3.name)
	fmt.Println(person3.age)
	fmt.Println()

	var p4 myStruct = myStruct{"park", 30}
	fmt.Println(p4)
	fmt.Println(p4.name)
	fmt.Println(p4.age)
}
