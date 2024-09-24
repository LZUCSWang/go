package main

import (
	"fmt"
)

var global *int

func main() {
	p := new(int)   // new返回的是指针
	fmt.Println(*p) // 默认0值
	*p = 2
	fmt.Println(*p)

	p1 := new(int)
	q1 := new(int)
	fmt.Println(p1 == q1) // false

	// new是预声明函数，不是关键字，可以被重新定义
	var new = 1
	fmt.Println(new)

	// p1 := new(int) // 此时new是变量
	// q1 := new(int)
	// fmt.Println(p1 == q1) // false

	f()
	fmt.Println(*global)
	fmt.Println(gcd(10, 5))

	fmt.Println(fbnc_digui(10))
	fmt.Println(fbnc(10))

	// v,ok=m[key]
	// v,ok=x.(T)
	// v,ok=<-ch

	// 隐式赋值
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals,medals[0],medals[1],medals[2])
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var t int
	return &t
}

func f() {
	var x int
	x = 1
	global = &x //x使用堆空间，在f结束后不会被释放
}

func g() {
	y := new(int)
	*y = 1 //y使用栈空间，在g结束后会被释放
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fbnc_digui(n int) int {
	if n < 2 {
		return n
	}
	return fbnc_digui(n-1) + fbnc_digui(n-2)
}

func fbnc(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
