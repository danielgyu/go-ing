package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "hello"
	hasPrefix := strings.HasPrefix(hello, "he")
	fmt.Println(hasPrefix)
}
