package main

import (
	"fmt"
)

func test(anyMap map[string]interface{}) {
	fmt.Println(anyMap["hello"])
}

func main() {
	m := map[string]int{"hello": 10}
	test(m)
}
