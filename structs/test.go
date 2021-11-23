package main

import "fmt"

type testStruct struct {
	GetWorkflow func(int) string
}

func main() {
	ts := testStruct{GetWorkflow: func(i int) string {
		return "hello world"
	}}

	fmt.Println(ts.GetWorkflow(1))

	var tt interface
	tt = map[string]int{"foo":10}
}
