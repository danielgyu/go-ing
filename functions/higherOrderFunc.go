package main

import "fmt"

type flt func(int) bool

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) (even, odd []int) {
	for _, s := range sl {
		if f(s) {
			even = append(even, s)
		} else {
			odd = append(odd, s)
		}
	}
	return
}

func add(x int, y int) int {
	return x + y
}

func nameAndSum(s string, f func(int, int) int) int {
	fmt.Println(s)
	return f(1, 100)
}

func main() {
	res := nameAndSum("Hello HigherOrder", add)
	fmt.Printf("result is %d\n", res)

	input := []int{1, 2, 3, 4, 5, 7}
	even, odd := filter(input, isEven)
	fmt.Println("filter:::")
	fmt.Println(even)
	fmt.Println(odd)
}
