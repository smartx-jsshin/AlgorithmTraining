package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num int
	if n, err := fmt.Scan(&num); n != 1 || err != nil {
		os.Exit(1)
	}

	for i := 0; i < num; i++ {
		var x1, y1, r1 float64
		var x2, y2, r2 float64
		fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2)

		// # of Pos: Infinity
		if x1 == x2 && y1 == y2 && r1 == r2 {
			fmt.Println(-1)
			continue
		}

		if r1 < r2 { // circle1 should be larger than circle2
			swap(&x1, &x2)
			swap(&y1, &y2)
			swap(&r1, &r2)
		}

		var xdist float64 = math.Abs(x2 - x1)
		var ydist float64 = math.Abs(y2 - y1)
		var diagdist float64 = math.Sqrt(xdist*xdist + ydist*ydist)

		if diagdist <= r1 { // Two Circles are overlapped
			// # of Pos: 2
			if diagdist+r2 < r1 {
				fmt.Println(0)
			} else if diagdist+r2 == r1 {
				fmt.Println(1)
			} else { // diagdist < r1+r2
				fmt.Println(2)
			}
		} else { // Two Circle are separated
			// # of Pos: 2
			if diagdist > r1+r2 {
				fmt.Println(0)
			} else if diagdist == r1+r2 {
				fmt.Println(1)
			} else { // diagdist < r1+r2
				fmt.Println(2)
			}
		}
	}
}

func swap(a *float64, b *float64) {
	var tmp float64
	tmp = *a
	*a = *b
	*b = tmp
}
