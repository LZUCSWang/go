package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	// fmt.Println(summer[:20])//超过边界
	endlessSummer := summer[:5]
	fmt.Println(summer)
	fmt.Println(endlessSummer)
	fmt.Println(len(summer), cap(summer))
	fmt.Println(summer[6]) // 报错
}
