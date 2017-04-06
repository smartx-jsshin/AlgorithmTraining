package main

import (
	"fmt"
)

var cnt0, cnt1 int

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		cnt0 = 0
		cnt1 = 0
		var n int
		fmt.Scan(&n)
		fibonacci(n)
		fmt.Println(cnt0, cnt1)
	}
}
func fibonacci(n int) int {
	if n == 0 {
		cnt0++
		return 0
	} else if n == 1 {
		cnt1++
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
