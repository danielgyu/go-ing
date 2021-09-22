package main

import "fmt"

func main() {
	var myMap1 map[string]int
	myMap1 = map[string]int{"foo": 100}
	fmt.Println(myMap1)

	var myMap2 = make(map[string]int)
	myMap2["bar"] = 200
	fmt.Println(myMap2)
}
