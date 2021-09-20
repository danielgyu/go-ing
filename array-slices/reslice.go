package main

import "fmt"

func main() {
	var ar = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	var slice = ar[2:4]
	fmt.Println(slice)
	slice = slice[0:6]
	fmt.Println(slice)
}
