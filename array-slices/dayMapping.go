package main

import "fmt"

func findName(dayMap map[int]string, n int) string {
	val, exist := dayMap[n]
	if exist {
		return val
	}
	return "Wrong Key"
}

func main() {
	dayMap := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	r1 := findName(dayMap, 1)
	r2 := findName(dayMap, 10)

	fmt.Println("r1:", r1)
	fmt.Println("r2:", r2)
}
