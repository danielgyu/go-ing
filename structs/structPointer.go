package main

import (
	"fmt"
	"unsafe"
)

type Vehicle struct {
	size  float64
	speed float64
	price float64
}

type Sedan struct {
	vehicle *Vehicle
	name    string
}

type Mini struct {
	vehicle Vehicle
	name    string
}

func main() {
	mini_size := unsafe.Sizeof(Mini{})
	sedan_size := unsafe.Sizeof(Sedan{})
	fmt.Println(mini_size)
	fmt.Println(sedan_size)
}
