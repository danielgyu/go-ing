package main

import "fmt"

func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		left := 0
		right := 1

		for j := i + 1; j < len(array); j++ {
			if array[left] > array[right] {
				array[left], array[right] = array[right], array[left]
			}
			left += 1
			right += 1
		}
	}
	return array
}

func main() {
	arr := []int{4, 5, 2, 1, 3}
	res := bubbleSort(arr)
	fmt.Println(res)
}
