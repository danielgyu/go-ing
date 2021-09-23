package main

import "fmt"

type Camera struct {
	pixel int
}

func (c Camera) setPixel(pixel int) {
	c.pixel = 30
}

type SmartPhone struct {
	Camera
}

func main() {
	c := new(Camera)
	fmt.Println(c)
	c.pixel = 30
	fmt.Println(c)
	c.setPixel(100)
	fmt.Println(c)
}
