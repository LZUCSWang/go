package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34
	// 新建空map的表达式:map[string]int{}
	fmt.Println(ages)

	ages["alice"] = 32
	fmt.Println(ages["alice"])

	delete(ages, "alice") //
	fmt.Println(ages)

	//键不在map中也可以操作
	delete(ages, "alice")
	delete(ages, "alices")

	fmt.Println(ages)
	fmt.Println(ages["bob"])
	ages["bob"] += 1
	fmt.Println(ages)

	// _ = &ages["bob"] // map元素不是一个变量，不能获取它的地址
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	var names []string

	for name := range ages { // 忽略的第二个是值
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	//通过指定一个slice的长度能更加高效
	namesSlice := make([]string, 0, len(ages)) // 创建了一个初始元素为空但是容量足够容纳ages map 所有键的slice
	fmt.Println(namesSlice)
	var age map[string]int
	fmt.Println(age == nil)
	fmt.Println(len(age) == 0)
}
