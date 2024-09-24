package main

import (
	"fmt"
	"os"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	// var name type = expression  type 和 expression不能同时省略
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
	// 变量声明，用var
	// go 的初始值都是0
	var s string
	fmt.Println(s)

	var i, j, k int
	var b, ff, ss = true, 2.3, "four" //允许不同的初始化类型
	fmt.Println(i, j, k, b, ff, ss)

	// 变量可以通过返回多个值的函数进行初始化
	var name = "test"
	var file, err = os.Open(name)
	fmt.Println(file, err)
	// 短变量声明，用:= 只能用在函数中
	ii := 100
	var boiling float64 = 100
	var names []string
	fmt.Println(ii, boiling, names)
	// names[0]="wang"
	names = append(names, "wang", "xian", "yi")
	fmt.Println(ii, boiling, names)

	i, j = j, i
}


