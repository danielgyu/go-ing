package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Languages []string `json:"languages"`
}

type Vehicle struct {
	Name  string
	Brand string
}

func main() {
	p := Person{"lee", 30, []string{"korean", "chinese"}}
	js, err1 := json.Marshal(p)
	if err1 != nil {
		fmt.Println("error:", err1)
		return
	}
	fmt.Println("person:", string(js))
	fmt.Println("type:", reflect.TypeOf(js))

	v := &Vehicle{"beetle", "volks"}
	js2, err2 := json.Marshal(v)
	if err2 != nil {
		fmt.Println("error:", err2)
		return
	}
	fmt.Println("vehicle:", string(js2))

	var up Person
	json.Unmarshal(js, &up)
	fmt.Println("decoded person:", up)

	var uv Vehicle
	json.Unmarshal(js2, &uv)
	fmt.Println("decoded vehicle:", uv)
}
