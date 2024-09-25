package main

import "fmt"

// 数组
func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}
	fmt.Println(a, b, c)
	q := [...]int{1, 2, 3, 4}
	fmt.Printf("%d %T\n", len(q), q)
	// q = [5]int{1}// 编译错误，不能[5]int -> [4]int
	// 数组长度必须是常量表达式，在编译时就要确定
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "e", GBP: "g", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	r := [...]int{99: -1}
	fmt.Println(r[99], r[0])

	aa := [2]int{1, 2}
	bb := [...]int{1, 2}
	cc := [2]int{1, 3}
	fmt.Println(aa == bb, aa == cc, bb == cc)
	dd := [...]int{1, 2, 3}
	// fmt.Println(aa == dd) //can't compare [2]int==[3]int
	// fmt.Println(aa!=dd)//cant't compare [2]int [3]int
	fmt.Println(dd)
	
}
