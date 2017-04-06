package main

import "fmt"

func main() {
	var instr string
	fmt.Scan(&instr)

	l := len(instr)

	var sum int
	for i := 0; i < l; i++ {
		cur_char := instr[i]
		cur_time := 2
		switch {
		case cur_char >= 'A' && cur_char <= 'C':
			cur_time += 1
		case cur_char >= 'D' && cur_char <= 'F':
			cur_time += 2
		case cur_char >= 'G' && cur_char <= 'I':
			cur_time += 3
		case cur_char >= 'J' && cur_char <= 'L':
			cur_time += 4
		case cur_char >= 'M' && cur_char <= 'O':
			cur_time += 5
		case cur_char >= 'P' && cur_char <= 'S':
			cur_time += 6
		case cur_char >= 'T' && cur_char <= 'V':
			cur_time += 7
		case cur_char >= 'W' && cur_char <= 'Z':
			cur_time += 8
		default:
			cur_time = 0
		}
		sum += cur_time
	}
	fmt.Println(sum)
}
