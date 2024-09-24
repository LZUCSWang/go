package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const pi = 3.14159

const (
	pi1 = 3.14
	pi2 = 3
)

const Ipv4Len = 4

const noDelay time.Duration = 0 //常量声明指定类型和值，可以省略类型，但是不能省略值
const timeout = 5 * time.Minute

const (
	A = 1
	B
	C = "yes"
	D
)

func parseIpv4(s string) string {
	var p [Ipv4Len]string //因为Ipv4Len是常量，所以可以用来指定数组类型的长度。
	// var len1 int64
	// var p1 [len1]string
	return p[0]
}

type Weekday int

const (
	Sunday Weekday = iota // 从0开始自动递增
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const ( // 随着iota的递增，每个常量的值都是1<<iota
	FlagUp        Flags = 1 << iota // 1<<0
	FlagBroadcast                   // 1<<1
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
	fmt.Println(comma("123354125"))
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2)
	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(intsToStringRune([]int{1, 2, 3}))
	fmt.Println(comma_buf("2123124234"))
	fmt.Println(comma_float("-12345.1243"))
	fmt.Println(comma_float("12345.1234"))
	fmt.Println(judge_yigou("abc", "bca"))
	fmt.Println(judge_yigou("abcc", "bca1"))

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	s = fmt.Sprintf("x = %d", x)
	fmt.Println(s)
	x1, err1 := strconv.Atoi("123")
	y2, err2 := strconv.ParseInt("123", 10, 64)
	fmt.Println(x1, err1, y2, err2)
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println(A, B, C, D)
}
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 去掉最后一个'/'之前的部分
	s = s[slash+1:]
	// 保留最后一个'.'之前的部分
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s // 返回s的拷贝
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func intsToStringRune(value []int) string {
	var buf bytes.Buffer
	buf.WriteRune('起')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune('终')
	return buf.String()
}

func comma_buf(s string) string {
	var buf bytes.Buffer
	l := len(s)
	plus := l % 3
	for i, v := range s {
		buf.WriteString(string(v))
		if (i+4-plus)%3 == 0 && i != len(s)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func comma_float(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}

	// l := len(s)
	dot := strings.Index(s, ".")
	if dot != -1 {
		inte := s[:dot]
		p := len(inte) % 3
		for i, v := range inte {
			buf.WriteString(string(v))
			if (i-p+4)%3 == 0 && i != len(inte)-1 {
				buf.WriteString(",")
			}
		}
	}
	buf.WriteByte('.')
	flo := s[dot+1:]
	// p := len(flo) % 3
	for i, v := range flo {
		buf.WriteString(string(v))
		if (i+1)%3 == 0 && i != len(flo)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

// 判断两个字符串是否构成同文异构（含有相同的字符但排列顺序不同）
func judge_yigou(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var c1, c2 [256]int
	for _, v := range s1 {
		c1[v]++
	}
	for _, v := range s2 {
		c2[v]++
	}
	for i := 0; i < 256; i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}
