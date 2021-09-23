package main

type employee struct {
	salary float64
}

func (this *employee) giveRaise(pct float64) {
	this.salary += this.salary * 0.1
}

func main() {
	e := new(employee)
	e.salary = 6000.0
}
