package main

import "fmt"

func main() {
	var multi [3][5]int
	fmt.Println(multi)
	r1 := [5]int{1, 2, 3, 4, 5}
	r2 := [5]int{6, 7, 8, 9, 10}
	r3 := [5]int{11, 12, 13, 14, 15}
	multi[0] = r1
	multi[1] = r2
	multi[2] = r3
	fmt.Println("multi:", multi)

	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9, 100, 100}
	fmt.Println(values)
	values = append(values, row1)
	fmt.Println(values)
	values = append(values, row2)
	fmt.Println(values)
	values = append(values, row3)
	fmt.Println(values)
}
