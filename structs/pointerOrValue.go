package main

import "fmt"

type Vehicle struct {
	topSpeed int
}

func (this Vehicle) setSpeed(speed int) {
	this.topSpeed = speed
}

func (this *Vehicle) changeSpeed(speed int) {
	this.topSpeed = speed
}

func main() {
	v1 := Vehicle{100}
	v1.setSpeed(200)
	fmt.Println(v1.topSpeed)
}
