package main

import "fmt"

type myInterface interface {
	Hello()
}

type myStruct struct {
	name string
}

func (s myStruct) Hello() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("vim-go")

	var ci myInterface
	var cs = myStruct{"foo"}

	ci = cs
	if t, ok := ci.(myStruct); ok {
		fmt.Println("yes:", t)
	}

	var ci2 myInterface
	var cs2 = myStruct{"bar"}

	if t, ok := ci2.(myStruct); ok {
		fmt.Println("yes!", t)
	} else {
		fmt.Println("no")
	}
}
