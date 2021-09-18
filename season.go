package main

import "fmt"

func season(month int) string {
	switch month {
	case 1, 2, 12:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Wrong number range"
	}
}

func main() {
	var month int = 11
	var season string = season(month)
	fmt.Printf("Season of the month is %s\n", season)
}
