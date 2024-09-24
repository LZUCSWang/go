package main

// echo1 输出其命令行参数
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, 世界!")
	var s, sep string
	// 参数列表从1开始，0是命令本身
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "*"
	}
	fmt.Println(s)

	// for 无限循环
	//for {
	//	fmt.Println("Hello, 世界!")
	//}

	// for while 循环
	// for condition { }

	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "*"
	}
	// _ 空白符，表示索引，arg 表示值
	fmt.Println(s)
	// go 中的变量声明（声明不使用会报错）
	news := ""
	var news1 string
	var news2 = ""
	var news3 string = ""
	fmt.Println(news, news1, news2, news3)

	// join
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[1:], "*"))

	// t1 输出命令名字
	fmt.Println(os.Args[0])
	// t2 输出参数的索引和值
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	// t3 比较Join和+的性能，使用time包
	// t3/

	
}
