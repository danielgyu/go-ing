package main

import "fmt"

const LIMIT = 4

type Stack struct {
	idx  int
	data [LIMIT]int
}

func (st *Stack) push(n int) {
	if st.idx < LIMIT {
		st.data[st.idx] = n
		st.idx++
	} else {
		fmt.Println("not enough spaces")
	}
}

func (st *Stack) pop() int {
	if st.idx > 0 {
		st.idx--
		element := st.data[st.idx]
		st.data[st.idx] = 0
		return element
	} else {
		fmt.Println("no element to pop")
		return -1
	}
}

func (st Stack) String() string {
	return fmt.Sprintf("[0:%d], [1:%d], [2:%d], [3:%d]", st.data[0], st.data[1], st.data[2], st.data[2])
}

func main() {
	s := new(Stack)
	s.push(1)
	s.push(2)
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
	s.pop()
	s.pop()
}
