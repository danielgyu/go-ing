package main

import (
	"errors"
	"fmt"
)

func testFunc(n int) (int, error) {
	customError := errors.New("custom error string")
	if n < 0 {
		return 0, customError
	}
	return 100, nil
}

func main() {
	if _, err := testFunc(-1); err != nil {
		fmt.Println(err)
	}
}
