package main

import "fmt"

func main() {
    var n int = 5
    fmt.Printf("%d is at memory location %p\n", n, &n)

    var intPointer *int
    intPointer = &n
    fmt.Println(&n)
    fmt.Println(intPointer)
    fmt.Println(*intPointer)
}
