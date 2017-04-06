package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		var mat = make([][]int, n, n)
		for value := range mat {
			mat[value] = make([]int, m)
		}
		fmt.Println(placeBridges(n, m, mat))
	}
}

func placeBridges(n int, m int, mat [][]int) int {
	var i int
	for i = 0; i < n; i++ {
		for j := i; j < m-(n-i-1); j++ {
			if i == 0 {
				mat[i][j] = 1
			} else {
				var sum int = 0
				for k := i - 1; k < j; k++ {
					sum += mat[i-1][k]
				}
				mat[i][j] = sum
			}
		}
	}
	var res int = 0
	for j := i - 1; j < m; j++ {
		res += mat[i-1][j]
	}

	return res
}
