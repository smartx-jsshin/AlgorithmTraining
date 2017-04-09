package main

import "fmt"

func main() {
	var flagmat []bool = make([]bool, 10001)
	var tmp int
	for i := 1; i <= 10000; i++ {
		tmp = dfunc(i)
		if tmp <= 10000 {
			flagmat[tmp] = true
		}
	}
	for i := 1; i <= 10000; i++ {
		if flagmat[i] == false {
			fmt.Println(i)
		}
	}
}

func dfunc(n int) int {
	var res int = n
	for true {
		if n == 10000 {
			res = res + n/10000
			n = n % 10000
		} else if n >= 1000 {
			res = res + n/1000
			n = n % 1000
		} else if n >= 100 {
			res = res + n/100
			n = n % 100
		} else if n >= 10 {
			res = res + n/10
			n = n % 10
		} else {
			res = res + n
			break
		}
	}
	return res
}
