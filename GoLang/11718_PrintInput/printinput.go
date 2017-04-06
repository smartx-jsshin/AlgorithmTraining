package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	test, _ := reader.ReadString('\n')
	fmt.Print(test)
}
