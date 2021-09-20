package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	num_of_copied := copy(sl_to, sl_from)
	fmt.Println("num_of_copied:", num_of_copied)
	fmt.Println("sl_to:", sl_to)
}
