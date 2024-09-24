package main

import (
	"flag"
	"fmt"
	"strings"
)

// 命令行输入go run echo4.go -s "*" -n true "wang" "xian" "yi"
// 看效果
var n = flag.Bool("n", false, "omit trailing newline") // 使用名字，默认值，输出信息
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Printf("separator is %s\n", *sep)
}
