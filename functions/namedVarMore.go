package main

import (
	"fmt"
	"math"
)

func f() (ret int) {
	defer func() {
		fmt.Printf("ret before inc: %d\n", ret)
		ret++
	}()
	return 10
}

func main() {
	fmt.Println(f())
	fmt.Println(math.Pow(math.Pow(4, 2), 3))
}
