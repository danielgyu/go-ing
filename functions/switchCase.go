package main

import "fmt"

func main() {
	var n int = 100
	switch n {
	case 98, 99:
		fmt.Println("Equal to 98 or 99")
	case 100:
		fallthrough
	case 101:
		fmt.Println("Fallthrough from 100")
		fallthrough
	default:
		fmt.Println("In default")
	}

	m := 101
	switch {
	case m > 105:
		fmt.Println("Bigger than 105")
	case m < 101:
		fmt.Println("Less than 101")
	}

	var num int
	switch x := 100 + 100; {
	case x > 300:
		num = 1
		fmt.Println(num)
	case x < 300:
		num = 2
		fmt.Println(num)
	}
}
