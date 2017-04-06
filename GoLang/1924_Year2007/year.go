package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var days int = 0
	for i := 1; i < x; i++ {
		if i == 2 {
			days = days + 28
		} else if i == 4 || i == 6 || i == 9 || i == 11 {
			days = days + 30
		} else {
			days = days + 31
		}
	}
	days = days + y

	switch days % 7 {
	case 1:
		fmt.Println("MON")
	case 2:
		fmt.Println("TUE")
	case 3:
		fmt.Println("WED")
	case 4:
		fmt.Println("THU")
	case 5:
		fmt.Println("FRI")
	case 6:
		fmt.Println("SAT")
	case 0:
		fmt.Println("SUN")
	default:
	}
}
