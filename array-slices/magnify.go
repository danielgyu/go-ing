package main

import "fmt"

func magnifySlice(s []int, factor int) (result []int) {
	result = make([]int, (len(s) * factor))
	copy(result, s)
	return
}

func main() {
	slice := []int{1, 2, 3}
	factor := 5

	res := magnifySlice(slice, factor)
	fmt.Println(res)
}
