package main

import "fmt"

func main() {
	var t, n, m, k int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m, &k)
		var farm = make([][]int, n, n)

		for i, _ := range farm {
			farm[i] = make([]int, m)
		}

		for j := 0; j < k; j++ {
			var x, y int
			fmt.Scan(&x, &y)
			farm[x][y] = 1
		}
		fmt.Println(countGroups(farm))
	}
}

func countGroups(mat [][]int) int {
	var maxrow int = len(mat)
	var maxcol int = len(mat[0])

	var queue = make([][2]int, 2500)
	var qflag = 0

	var numGroups int = 0

	for i := 0; i < maxrow; i++ {
		for j := 0; j < maxcol; j++ {
			if mat[i][j] == 0 || mat[i][j] == 2 {
				continue
			} else if mat[i][j] == 1 {
				push(queue, &qflag, [2]int{i, j})
				numGroups += 1
			}

			for qflag > 0 {
				var v = pop(queue, &qflag)
				if mat[v[0]][v[1]] != 2 {
					mat[v[0]][v[1]] = 2
					// Below Node
					if v[0]+1 < maxrow && mat[v[0]+1][v[1]] == 1 {
						push(queue, &qflag, [2]int{v[0] + 1, v[1]})
					}
					// Right Node
					if v[1]+1 < maxcol && mat[v[0]][v[1]+1] == 1 {
						push(queue, &qflag, [2]int{v[0], v[1] + 1})
					}
					// Up Node
					if v[0]-1 >= 0 && mat[v[0]-1][v[1]] == 1 {
						push(queue, &qflag, [2]int{v[0] - 1, v[1]})
					}
					// Left Node
					if v[1]-1 >= 0 && mat[v[0]][v[1]-1] == 1 {
						push(queue, &qflag, [2]int{v[0], v[1] - 1})
					}
				}
			}
		}
	}
	return numGroups
}

func pop(queue [][2]int, qflag *int) [2]int {
	*qflag -= 1
	var tmp [2]int = queue[*qflag]
	return tmp
}

func push(queue [][2]int, qflag *int, v [2]int) {
	queue[*qflag] = v
	*qflag += 1
}
