package main

import "fmt"

func main() {
	//var array = [5]int{1, 2, 3, 4, 5}
	//var slice = array[:]

	//var slice = make([]int, 5)

	var array2 = new([5]int)
	var slice = array2[:]

	for i := range slice {
		if i == 0 || i == 1 {
			slice[i] = 1
			continue
		}

		slice[i] = slice[i-1] + slice[i-2]
	}

	fmt.Println(slice)
}
