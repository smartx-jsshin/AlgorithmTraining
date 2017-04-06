package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var instr string
	var cnt int

	for i := 0; i < n; i++ {
		fmt.Scan(&instr)

		var prev_char byte
		var check_mat [26]bool
		group_word := true

		for j := 0; j < len(instr); j++ {
			idx := int(instr[j] - 'a')
			if prev_char != instr[j] {
				if check_mat[idx] == false {
					check_mat[idx] = true
				} else {
					group_word = false
					break
				}
			}
			prev_char = instr[j]
		}

		if group_word == true {
			cnt = cnt + 1
		}
	}
	fmt.Println(cnt)
}
