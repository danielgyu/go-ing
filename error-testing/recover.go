package main

import "fmt"

func badFunction() {
	panic("bad function")
}

func testFunction() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovering from", e)
		}
	}()
	fmt.Println(recover())
	badFunction()
}

func main() {
	testFunction()
}
