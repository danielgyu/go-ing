package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main()")
	go longWait()
	go shortWait()
	fmt.Println("beginning sleep in main()...")
	time.Sleep(7 * 1e9)
	fmt.Println("end of main()")
}

func longWait() {
	fmt.Println("beginning longWait()...")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longwait()")
}

func shortWait() {
	fmt.Println("beginning shortWait()...")
	time.Sleep(2 * 1e9)
	fmt.Println("end of shortWait()")
}
