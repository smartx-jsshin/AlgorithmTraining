package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 2*n; i++ {
		var starcnt int
		if i > n {
			starcnt = n - (i - n)
		} else {
			starcnt = i
		}

		for j := 1; j <= starcnt; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= (2*n - 2*starcnt); j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= starcnt; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
