package main

import "fmt"

func main() {
	var arr [10]int

	for i := range arr {
		if i == 0 || i == 1 {
			arr[i] = 1
			continue
		}

		arr[i] = arr[i-1] + arr[i-2]
	}

	fmt.Println(arr)
}
