package main

import (
	"fmt"
)

func main() {
	var cnt_arr [26]int
	var instr string

	fmt.Scanf("%s", &instr)
	l := len(instr)
	idx := -1
	for i := 0; i < l; i++ {
		if instr[i] >= 'a' && instr[i] <= 'z' {
			idx = int(instr[i] - 'a')
		} else if instr[i] >= 'A' && instr[i] <= 'Z' {
			idx = int(instr[i] - 'A')
		}
		cnt_arr[idx] += 1
	}

	idx = 0
	dup := false
	for i := 1; i < 26; i++ {
		if cnt_arr[idx] < cnt_arr[i] {
			idx = i
			dup = false
		} else if cnt_arr[idx] == cnt_arr[i] {
			dup = true
		}
	}

	if dup == true {
		fmt.Println("?")
	} else {
		fmt.Printf("%c", byte(idx)+'A')
	}
}
