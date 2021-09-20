package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s)
	fmt.Println(c)

	c[0] = 'H'
	s2 := string(c)
	fmt.Println(s2)
}
