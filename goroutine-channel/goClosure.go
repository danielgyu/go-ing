package main

import "fmt"

func main() {
	data := []int{1, 2, 3}

	for idx, v := range data {
		func() {
			v += 1
			fmt.Println("idx:", idx)
			fmt.Println("v:", v)
		}()
	}

	fmt.Println("data:", data)
}
