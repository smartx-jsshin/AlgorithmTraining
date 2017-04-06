package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		var mat = make([][]int, n)
		for j := 0; j < n; j++ {
			mat[j] = make([]int, m)

			var tmpstr string
			fmt.Scan(&tmpstr)
			initMatRow(mat[j], tmpstr)
		}
		for _, val := range mat {
			fmt.Println(val)
		}
		fmt.Println("")
	}
}

func initMatRow(row []int, str string) {
	for pos, val := range str {
		if val == 'x' || val == 'X' {
			row[pos] = 1
		}
	}
}
