// slice就地修改算法
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	strings := []string{"", "Hello", "Wang", ""}
	fmt.Println(strings)
	newStrings := nonempty(strings)
	fmt.Println(strings, newStrings)
	fmt.Println(len(strings), len(newStrings))
	newStrings2 := nonempty2(strings)
	fmt.Println(len(strings), len(newStrings2))
	fmt.Println(strings, newStrings2)
	// slice可以实现栈
	stack := []int{}
	v := 1
	stack = append(stack, v)
	s := []int{5, 6, 7, 8, 9}
	s = remove(s, 3)
	fmt.Println(s)
	arry := [5]int{1, 2, 3, 4, 5}
	reverse_arry(&arry)
	fmt.Println(arry)
	reverse_arry(&arry)

	arry_rightShifed := rotate(arry[:], 2)
	fmt.Println(arry_rightShifed)

	arry_str := []string{"H", "E", "L", "L", "O"}
	arry_str = removeAdjacentDuplicateStrings(arry_str)
	fmt.Println(arry_str)
	slice := []byte("Hello,\t\tworld!\n\nThis is a test.\r\n")
	fmt.Println("原始切片:", string(slice))
	slice = reduceAdjacentWhitespace(slice)
	fmt.Println("处理后的切片:", string(slice))
}
func remove_without_order(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
func push(stack []int, v int) []int {
	stack = append(stack, v)
	return stack
}

func pop(stack []int, v int) int {
	top := stack[len(stack)]
	stack = stack[:len(stack)-1]
	return top
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 用append实现nonempty
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		out = append(out, s)
	}
	return out
}

// 重写函数reverse，使用数组指针作为参考而不是slice
func reverse_arry(arry *[5]int) {
	for i, j := 0, len(arry)-1; i < j; i, j = i+1, j-1 {
		arry[i], arry[j] = arry[j], arry[i]
	}
}

// 编写一个函数rotate，实现一次遍历就可以完成元素旋转
func rotate(slice []int, k int) []int {
	n := len(slice)
	k %= n
	reverse_rotate(slice, 0, n-1)
	reverse_rotate(slice, 0, k-1)
	reverse_rotate(slice, k, n-1)
	return slice
}

func reverse_rotate(slice []int, start int, end int) {
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
}

// 编写一个就地处理函数，用于去除[]string slice 中相邻的重复字符串元素
func removeAdjacentDuplicateStrings(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	j := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[j] {
			j++
			slice[j] = slice[i]
		}
	}
	return slice[:j+1]
}

// 编写一个就地处理函数，用于将一个UTF-8编码的字节slice中所有相邻
// 的unicode空白字符缩减为一个ASCII空白字符

func reduceAdjacentWhitespace(slice []byte) []byte {
	if len(slice) == 0 {
		return slice
	}
	j := 0
	spaceFound := false
	for i := 0; i < len(slice); {
		r, size := utf8.DecodeLastRune(slice[i:])
		if unicode.IsSpace(r) {
			if !spaceFound {
				slice[j] = ' '
				j++
				spaceFound = true
			}
		} else {
			copy(slice[j:], slice[i:i+size])
			j += size
			spaceFound = false
		}
		i += size
	}
	return slice[:j]
}
