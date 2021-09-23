package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p *Point) Scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
	return
}

func main() {
	p := new(Point)
	p.x = 3
	p.y = 4
	fmt.Println(p.Abs())
	p.Scale(5)
	fmt.Println(p.Abs())

	p1 := Point{9, 10}
	fmt.Println(p1.Abs())
}
