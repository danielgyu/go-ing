package main

import "fmt"

func insertSlice(slice, insertion []string, index int) []string {
	var result []string

	for i, s := range slice {
		if i == index {
			for j := range insertion {
				result = append(result, insertion[j])
			}
		}
		result = append(result, s)
	}

	return result
}

func main() {
	slice := []string{"M", "N", "O", "P"}
	insertion := []string{"A", "B", "C"}
	index := 0

	res := insertSlice(slice, insertion, index)
	fmt.Println(res)
}
