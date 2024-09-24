package main

import (
	"fmt"
)

func main() {
	fmt.Println((!true == false) == true)
	s := ""
	// fmt.Println(s == "" && s[0] == 'x') //虽然短路，但是仍然会报错
	fmt.Println(s == "")
	// i, b := 0, 1
	i, b := 0, true
	// i, b := 0, 1 // go不支持int和bool的隐式转换
	if b {
		i = 1
	}
	fmt.Println(i)
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
