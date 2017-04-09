package main

import "fmt"

func main() {
	var initval, tmpval, cnt int
	fmt.Scan(&initval)

	tmpval = initval
	cnt = 0
	for true {
		var a, b int
		a = tmpval % 10
		b = (tmpval/10 + tmpval%10) % 10
		tmpval = a*10 + b
		cnt = cnt + 1

		//fmt.Println("cnt: ", cnt, " new: ", tmpval)
		if tmpval == initval {
			break
		}
	}
	fmt.Println(cnt)
}
