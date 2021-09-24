package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Number struct {
	X, Y int
}

func main() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(Number{3, 4})
	if err != nil {
		return
	}
	fmt.Println("after encode, buffer:", buf)

	var decodedNum Number
	dec := gob.NewDecoder(&buf)
	err2 := dec.Decode(&decodedNum)
	if err2 != nil {
		return
	}
	fmt.Println("decoded:", decodedNum)
	fmt.Println(decodedNum.X)
}
