package main

import "fmt"

type Vehicle struct {
	name string
	age  int
}

type Automobile Vehicle

func (recv Vehicle) printName() {
	fmt.Println(recv.name)
}

func (recv Vehicle) getAge() int {
	return recv.age
}

func (_ Vehicle) greet() {
	fmt.Println("hello!")
}

func (_ *Vehicle) honk() {
	fmt.Println("honk")
}

func main() {
	v := Vehicle{"my vehicle", 100}
	fmt.Println(v)
	v.printName()
	fmt.Println(v.getAge())
	v.greet()
	v.honk()
}
