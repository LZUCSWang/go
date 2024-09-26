package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v", i, cap(y), y)
		fmt.Println(x) // appendInt 不会修改原来的slice，只会生成新的slice
		x = y
	}

	// 内置的append函数使用了更复杂的增长策略，不能假设原始的和新的会不会相互影响
	// 故 runes = append(runes,r)比较保险

	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
	z := []int{1, 2, 3}
	z = appendInt1(z, z...)
	fmt.Println(z)

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		// 对slice进行扩容
		z = x[:zlen]
	} else {
		// zlen无空间，分配一个新的底层数组
		zcap := zlen
		if zcap < 2*len(x) { //指数增长，未来的分配时间会更少
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInt1(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen < cap(x) {
		// 对slice进行扩容
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
