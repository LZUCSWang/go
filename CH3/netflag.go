package main

import (
	"fmt"
	"math"
)

const Pi64 float64 = math.Pi

var x float32 = float32(Pi64)
var y float64 = Pi64
var z complex128 = complex128(Pi64) // 需要用到类型转换

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBoradcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

const (
	deadbeef = 0xdeadbeef
	a        = uint32(deadbeef)
	b        = float32(deadbeef)
	c        = float64(deadbeef)
	// d        = int32(deadbeef) // 编译错误:溢出，int32无法容纳常量值
	// e        = float64(1e309)  // 编译错误:溢出
	// f        = uint(-1)        // 编译错误:溢出
)

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) //"10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) //"10000 false"
	SetBoradcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) //"10010 true"

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100",(f-32)*5的结果是float64型
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9的结果是无类型整数0；
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100";5.0/9.0 的结果是无类型浮点数
	fmt.Println(5 / 9.0 * (f - 32))

	// 只有常量才可以是无类型的。
	var f1 float64 = 3 + 0i //无类型复数 ——> float64
	f1 = 2                  // 无类型整数 -> float64
	f1 = 1e123              // 无类型浮点数 -> float64
	f1 = 'a'                // 无类型->float64
	fmt.Println(f1)
	// 上面的和下面的等效
	var f2 float64 = float64(3 + 0i)
	f2 = float64(2)
	f2 = float64(1e123)
	f2 = float64('a')
	fmt.Println(f2)

	// 如果没有显示指定类型，无类型常量会隐式转换成该变量的默认类型
	i := 0      // 无类型整数，隐式int(0)
	r := '\000' // 无类型文字字符，隐式rune('\000')
	ff := 0.0   // 无类型浮点数，隐式float64(0.0)
	c := 0i     // 无类型整数
	// 将无类型转为接口值时，默认类型就很重要，决定接口值的动态类型
	fmt.Println(i, r, ff, c)
	fmt.Printf("%T\n%T\n%T\n%T\n", 0, 0.0, 0i, '\000')
}
