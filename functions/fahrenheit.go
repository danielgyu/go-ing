package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func celsiusToFahrenheit(celsius Celsius) Fahrenheit {
	return Fahrenheit(celsius*(9.0/5.0) + 32.0)
}

func main() {
	var celsius Celsius = 0.0
	fahrenheit := celsiusToFahrenheit(celsius)

	fmt.Println(fahrenheit)
}
