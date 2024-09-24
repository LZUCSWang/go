package tempconv
// package main
// import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func cToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func fToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 这是一个包，相当于库函数