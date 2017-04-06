package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	ares := getReverseInt(a)
	bres := getReverseInt(b)

	if ares >= bres {
		fmt.Println(ares)
	} else {
		fmt.Println(bres)
	}
}

func getReverseInt(n string) int {
	var res int
	for i := len(n) - 1; i >= 0; i-- {
		res = res + int(n[i]-'0')*int(math.Pow10(i))
	}
	return res
}
