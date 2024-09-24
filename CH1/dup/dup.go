package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// dup1
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 { // if 不用括号
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
