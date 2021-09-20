package main

import "fmt"

func copyAndAdd(arr [3]int) [3]int {
	for i := range arr {
		arr[i] += 1
	}
	return arr
}

func addOne(arr *[4]int) *[4]int {
	for i := range *arr {
		arr[i] += 1
	}
	return arr
}

func main() {
	var arr1 = [4]int{1, 2, 3}
	fmt.Println(arr1)

	addOne(&arr1)
	fmt.Println("After adding:", arr1)

	var arr2 = [3]string{1: "Lee", 2: "Kim"}
	fmt.Println(arr2[0])
	fmt.Println(arr2)

	var arr3 = [3]int{10, 11, 12}
	fmt.Println(arr3)
	copyAndAdd(arr3)
	fmt.Println(arr3)
}
