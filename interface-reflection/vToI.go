package main

import "fmt"

type Stringer interface {
	String() string
}

type Keyboard struct {
	key_nums int
}

func (k Keyboard) shout() string {
	return "from keyboard"
}

func (k Keyboard) String() string {
	return fmt.Sprintf(k.shout())
}

func main() {
	fmt.Println("vim-go")

	v := Keyboard{10}
	v1 := new(Keyboard)
	v1.key_nums = 100
	if kv, ok := v1.(Stringer); ok {
		fmt.Println("finished", kv)
	}
}
