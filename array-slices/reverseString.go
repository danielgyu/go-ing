package main

import "fmt"

func reverse(s string) (resultString string) {
	b := []byte(s)
	res := make([]byte, len(b))

	for i := 0; i < len(b); i++ {
		res[i] = b[len(b)-i-1]
	}

	resultString = string(res)
	return
}

func main() {
	res := reverse("Google")
	fmt.Println("final:", res)
}
