package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonData = []byte(`{"name": "lee", "age": 30, "time": 92.30}`)

	var decodeAnything interface{}

	err := json.Unmarshal(jsonData, &decodeAnything)
	if err != nil {
		return
	}
	fmt.Println("decoded:", decodeAnything)
}
