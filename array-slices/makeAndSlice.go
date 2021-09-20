package main

import "fmt"

func main() {
	newArray := new([10]int)
	makeArray := make([]int, 3, 5)
	fmt.Println(newArray)
	fmt.Println(makeArray)

	fmt.Printf("length of newArray: %d\n", len(*newArray))
	fmt.Printf("cap of newArray: %d\n", cap(*newArray))

}
