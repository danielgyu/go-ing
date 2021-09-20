package main

import "fmt"

func main() {
	var array1 = new([5]int)
	var array2 [5]int

	fmt.Println(array1)
	fmt.Println(array2)

	var array3 *[5]int = array1
	var array4 [5]int = array2

	array3[0] = 1000
	array4[0] = 200

	fmt.Println(array1)
	fmt.Println(array2)
}
