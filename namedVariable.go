package main

import "fmt"

func addTwoNumbers(n int, m int) (s int) {
	s = n + m
	return
}

func sumProductDiff(i, j int) (sum int, product int, difference int) {
	sum = i + j
	product = i * j
	difference = i - j
	return
}

func sumInts(nums ...int) (sum int) {
	sum = 0
	fmt.Println(nums)
	for _, n := range nums {
		fmt.Printf("n is %d\n", n)
		sum += n
	}
	return
}

func main() {
	sum := addTwoNumbers(10, 20)
	fmt.Println(sum)

	s, p, d := sumProductDiff(3, 4)
	fmt.Printf("%d %d %d\n", s, p, d)

	sumints := sumInts(1, 2, 3)
	fmt.Println(sumints)
}
