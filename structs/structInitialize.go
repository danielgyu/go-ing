package main

import "fmt"

type Point struct{ x, y int }

func main() {
	pp := &Point{30, 40}

	fmt.Println(pp.x, pp.y)

	ppp := &Point{50, 60}
	fmt.Println(pp.x, pp.y)
	fmt.Println(ppp.x, ppp.y)
}
