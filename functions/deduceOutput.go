package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func printThat(s string) string {
	fmt.Println("print that", s)
	return s
}

func c() {
	fmt.Println("c")
	defer un(printThat("c"))
}

func d() {
	fmt.Println("d")
	defer printThat("d")
}

func main() {
	//b()
	c()
	//d()
}
