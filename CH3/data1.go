package main

import (
	"fmt"
)

func main() {
	f := 3.141
	i := int(f)
	fmt.Println(i, f)
	f = 1.99
	fmt.Println(int(f)) // 去掉小数部分

	fl := 1e100   // float64
	fi := int(fl) // 会导致溢出，超出原来的范围
	fmt.Println(fl, fi)

	o := 0666 //八进制
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// [1]x 表示使用第一个参数，%[1]x 表示使用第一个参数的十六进制格式

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
	fmt.Printf("%c\n", newline)
}
