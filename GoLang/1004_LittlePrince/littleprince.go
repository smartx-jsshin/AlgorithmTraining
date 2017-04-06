package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		var x1, y1, x2, y2 float64
		fmt.Scan(&x1, &y1, &x2, &y2)

		var planum int
		fmt.Scan(&planum)

		var cntcross int = 0
		for j := 0; j < planum; j++ {
			var cx, cy, r float64
			var isD1inC, isD2inC bool
			fmt.Scan(&cx, &cy, &r)

			isD1inC = checkDotInCircle(x1, y1, cx, cy, r)
			isD2inC = checkDotInCircle(x2, y2, cx, cy, r)

			if isD1inC && !isD2inC || !isD1inC && isD2inC {
				cntcross++
			}
		}

		fmt.Println(cntcross)
	}
}

func checkDotInCircle(x float64, y float64, cx float64, cy float64, r float64) bool {
	var tdist float64 = math.Sqrt((cx-x)*(cx-x) + (cy-y)*(cy-y))
	if tdist < r {
		return true
	} else {
		return false
	}
}
