package main

import "fmt"

type Stack []interface{}

func (st Stack) Length() int {
	return len(st)
}

func (st Stack) isEmpty() bool {
	return len(st) == 0
}

func (st *Stack) Push(n interface{}) {
	*st = append(*st, n)
}

func (st *Stack) Pop() (x interface{}, err error) {
	if st.isEmpty() {
		return -1, err
	}
	idx := st.Length() - 1
	*st = (*st)[:idx]
	return 1, nil
}

func main() {
	var st Stack
	fmt.Println(st.Length())
	fmt.Println(st.isEmpty())

	st.Push(8)
	st.Push(8.8)
	st.Push("foo")
	st.Push(map[string]int{"bar": 100})
	fmt.Println(st)

	st.Pop()
	fmt.Println(st)
}
