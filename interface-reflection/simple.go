package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	year int
}

func (s Simple) Get() int {
	return s.year
}

func (s *Simple) Set(n int) {
	s.year = n
}

type RSimple struct {
	year int
}

func (rs RSimple) Get() int {
	return rs.year
}

func (rs *RSimple) Set(n int) {
	rs.year = n
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		fmt.Println("Simple")
		return it.Get()
	case *RSimple:
		fmt.Println("RSimple")
		return it.Get()
	default:
		return -1
	}
}

func interfaceImplementation(si Simple) bool {
	if _, ok := si.(Simpler); ok {
		return true
	} else {
		return false
	}
}

func main() {
	s := new(Simple)
	s.Set(1992)
	fmt.Println(s.Get())

	rs := new(RSimple)
	rs.Set(2002)

	fmt.Println(fI(s))
	fmt.Println(fI(rs))

	fmt.Println(interfaceImplementation(*s))
}
