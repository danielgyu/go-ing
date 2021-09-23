package main

import "fmt"

type Element interface{}

type Vector struct {
	elem []Element
}

func (v *Vector) At(i int) Element {
	return v.elem[i]
}

func (v *Vector) Set(i int, e Element) {
	v.elem[i] = e
}

func main() {
	fmt.Println("vim-go")
}
