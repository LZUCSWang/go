// 字符串是不可变的字节序列
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Printf("%c %c\n", s[0], s[7])
	fmt.Println("goodbye" + s[5:])
	s = "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s)
	fmt.Println(t)
	// s[0] = 'L' // cannot assign to s[0]

	s = "Hello, 世界"
	fmt.Println(len(s))                    // 13 按照字节长度计算，一个汉字占3个字节
	fmt.Println(utf8.RuneCountInString(s)) // 9 按照rune长度计算，一个汉字占1个rune
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) // 返回rune（文字符号本身）和rune的长度
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	for range "Hello, 世界" {
		n++
	}
	fmt.Println(n)
}
