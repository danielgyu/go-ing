package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{1, 2, 3}
	var slice1 []int = arr[:]
	fmt.Println(slice1)
	fmt.Println()

	var arr2 = [6]int{1, 2, 3, 4, 5, 6}
	var slice2 []int = arr2[2:4]
	fmt.Println(slice2)
	slice2 = slice2[0:4]
	fmt.Println(slice2)
	fmt.Println()

	var slice3 []int = make([]int, 3, 10)
	for i := range slice3 {
		slice3[i] += 10
	}
	slice3[3] = 100
	for i := range slice3 {
		fmt.Println(slice3[i])
	}
}
