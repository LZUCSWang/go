package main

import (
	"fmt"
	"log"
	"os"
)

// var g = "g"
var golbal int

func main() {
	// f := "f"
	// fmt.Println(f)
	// fmt.Println(g)
	// // fmt.Println(h)
	global := 1
	fmt.Println(global) // 此处的global是局部变量，并不是对全局变量的引用
	glo()               //此处的glo函数中的golbal是全局变量
	x := "htllo!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()

	// x := "hello!"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println()

	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == 1 {
		fmt.Println(x, y)
	}

}
func glo() {
	fmt.Println(golbal)
}
func f() int {
	fmt.Printf("f")
	return 0
}
func g(int) int {
	fmt.Printf("g")
	return 1
}

func openfile() {
	fname := "example.txt" // Define the fname variable
	if f, err := os.Open(fname); err != nil {
		fmt.Println(err) // Print the error instead of returning it
		return
	} else {
		f.Stat()
		f.Close()
	}
}

var cwd string

func init() {
	cwd, err := os.Getwd() //:=会将cwd变量重新定义为局部变量，而不是赋值给全局变量
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
		fmt.Println(cwd)
	}
}

// 上面的不好，应该将err定义为新的变量，而不是使用:=

func init1() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
		fmt.Println(cwd)
	}
}
