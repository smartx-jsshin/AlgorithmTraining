package main

import "fmt"

func main() {
	var n, cnt int
	cnt = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if hanSu(i) {
			cnt = cnt + 1
		}
	}
	fmt.Println(cnt)
}

func hanSu(n int) bool {
	var digit []int

	switch {
	case n == 1000:
		digit = make([]int, 4)
	case n >= 100:
		digit = make([]int, 3)
	case n >= 10:
		digit = make([]int, 2)
	case n >= 1:
		digit = make([]int, 1)
	default:
		digit = nil
	}

	for true {
		if n == 1000 {
			digit[3] = n / 1000
			n = n % 1000
		} else if n >= 100 {
			digit[2] = n / 100
			n = n % 100
		} else if n >= 10 {
			digit[1] = n / 10
			n = n % 10
		} else {
			digit[0] = n
			break
		}
	}

	var gap int = -255
	for i := 0; i < len(digit)-1; i++ {
		if gap != -255 {
			//fmt.Println("gap: ", gap, " cur: ", digit[i]-digit[i+1])
			if gap != (digit[i] - digit[i+1]) {
				return false
			}
		}
		gap = digit[i] - digit[i+1]
	}
	return true
}
