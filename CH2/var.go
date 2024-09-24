package main

import (
	"fmt"
	"os"
)

func main() {
	//:=是声明 =是赋值 短 变量声明不需 要声明 所有在左边的 变量。如果一些
	// 变量在 同一个 词法块 中声明 (参考 2.7 节)，那么对于那些 变量，短声明的行为等同于赋值
	infile, outfile := "in", "out"
	in, err := os.Open(infile)
	//...
	out, err := os.Create(outfile) //此处编译通过，因为这句中有一个新的变量
	fmt.Println(in, err, out)
	// in,err :=os.Open(infile) 会报错，因为这个声明语句中没有新的变量

	// 也就是说至少要有一个新的变量声明语句才不会报错，老变量就当作是复制
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	var q *int //指针的0值是nil，不会像c一样乱指
	fmt.Println(q, q != nil)

	var x1, y1 int
	fmt.Println(&x1 == &x1, &y1 == &y1, &x1 == nil)

	var fp = f()
	fmt.Println(f() == f(), fp) // 每次调用f()都会返回不同的指针值

	y := 1
	fmt.Println(incr(&y), y)

}

func f() *int {
	var x = 1
	return &x
}

func incr(p *int) int {
	*p++
	return *p
}
