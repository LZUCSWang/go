package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	a := [...]int{1, 3, 5, 7, 9}
	reverse(a[:]) //一个数组的slice就相当于数组的一个别名，可以通过它
	// 修改数组
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	// 将a左移2个元素[5,7,9,1,3]
	reverse(a[:2]) // [3,1,5,7,9]
	reverse(a[2:]) // [3,1,9,7,5]
	reverse(a[:])  // [5,7,9,1,3]
	fmt.Println(a)

	//slice 无法进行比较
	sa, sb := []string{"he"}, []string{"ll"}
	fmt.Println(equal(sa, sb))
	if sa == nil {
		fmt.Println("summer is nil")
	}
	var varsnil []int
	if varsnil == nil {
		fmt.Println("varsnil is nil")
	}
	snil := []int{}
	if snil == nil {
		fmt.Println("snil is nil")
	}
	var s1 []int //len(s) == 0,s==nil
	fmt.Println(len(s1), s1 == nil)
	s1 = nil //len(s) == 0, s1 == nil
	fmt.Println(len(s1), s1 == nil)
	s1 = []int(nil) // len(s) == 0, s2 == nil
	fmt.Println(len(s1), s1 == nil)
	s1 = []int{} //len(s) == 0, s != nil
	fmt.Println(len(s1), s1 == nil)
	//检查slice是否为空用len不用==nil

	//用make创建slice
	s2 := make([]int, 10, 50)
	fmt.Println(len(s2), cap(s2))
	// make([]T,len,cap)创建一个数组切片，len是数组切片的长度，cap是数组切片的容量
	// make ([]T,cap)[:len]创建一个长度为len的数组切片
	// make([]T,len,cap)和make([cap]T)[:len]是等价的
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
