package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Println("Using runtime::")
	fmt.Println(runtime.Caller(1))
	fmt.Println()

	fmt.Println("Using log")
	log.SetFlags(log.Llongfile)
	log.Print("Flag set")
}
