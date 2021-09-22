package main

import "fmt"

func main() {
	var myMap = map[string]int{"foo": 0}
	v1, isPresent1 := myMap["foo"]
	v2, isPresent2 := myMap["boo"]
	fmt.Println(v1, isPresent1)
	fmt.Println(v2, isPresent2)

	fmt.Println("Before delete:", myMap)
	delete(myMap, "foo")
	fmt.Println("After delete:", myMap)
}
