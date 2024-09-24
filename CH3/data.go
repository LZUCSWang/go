package main

import (
	"fmt"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var j uint8 = 0
	// ^ bitwise NOT, ^在一元前缀运算符中表示按位取反
	fmt.Printf("%08b\n", ^u)
	fmt.Printf("%08b\n", ^i)
	fmt.Printf("%08b\n", ^j)

	// &^ bitwise AND NOT, &^在二元中表示按位清零
	fmt.Printf("%08b\n", u&^0)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- { //此处想说明的是len返回的是int类型，所以i也是int类型，否则i会变成uint的最大值导致越界访问元素
		fmt.Println(medals[i])
	}

	var apple int32 = 1
	var oranges int16 = 2
	// var compote int = apple + oranges //编译错误
	var compote int = int(apple) + int(oranges)
	fmt.Println(compote)
}
