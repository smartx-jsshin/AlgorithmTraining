package main

import "fmt"
import "strings"

var mat [][]string

func main() {
	var n int
	fmt.Scan(&n)

	tmp := make([]string, n*2-1)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = " "
	}

	mat = make([][]string, n)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]string, n*2-1)
		copy(mat[i], tmp)
	}

	drawstar(n, n-1, 0)
	visualize(n, n*2-1)
}

func drawstar(n int, center int, top int) {
	//fmt.Println("drawstar", n, center, top)

	if n == 3 {
		//  *
		mat[top][center] = "*"
		// * *
		mat[top+1][center-1] = "*"
		mat[top+1][center+1] = "*"
		//*****
		mat[top+2][center-2] = "*"
		mat[top+2][center-1] = "*"
		mat[top+2][center] = "*"
		mat[top+2][center+1] = "*"
		mat[top+2][center+2] = "*"

	} else {
		drawstar(n/2, center, top)
		drawstar(n/2, center-n/2, top+n/2)
		drawstar(n/2, center+n/2, top+n/2)
	}
}

func visualize(height int, width int) {
	for i := 0; i < height; i++ {
		fmt.Println(strings.Join(mat[i], ""))
	}
}
