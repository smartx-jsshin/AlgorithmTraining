package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, res int
		var c int = 1
		fmt.Scan(&a, &b)
		for j := 0; j < b; j++ {
			c = c * a
			if c > math.MaxInt8 {
				c = c % 10
			}
		}
		res = c % 10
		if res == 0 {
			fmt.Println("10")
		} else {
			fmt.Println(res)
		}
	}
}
