package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	fmt.Println(b)

	b.WriteString("ABC")
	fmt.Println(b)

	fmt.Fprintf(&b, "number: %d\n", 10)
	fmt.Println(b)
}
