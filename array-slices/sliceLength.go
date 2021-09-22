package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr = arr[2:2]
	fmt.Println(arr)
}
