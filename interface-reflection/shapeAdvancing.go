package main

import "fmt"

type AreaInterface interface {
	getArea() float32
}

type Triangle struct {
	base   float32
	height float32
}

func (this Triangle) getArea() float32 {
	return 0.5 * (this.base * this.height)
}

type PerimeterInterface interface {
	getPerimeter()
}

type Square struct {
	side float32
}

func (this Square) getArea() float32 {
	return this.side * this.side
}

func (this Square) getPerimeter() float32 {
	return 4 * this.side
}

func main() {
	fmt.Println("vim-go")
}
