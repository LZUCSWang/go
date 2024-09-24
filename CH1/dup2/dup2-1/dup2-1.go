package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0{
		countLines(os.Stdin,counts)
	}else{
		for _,arg:=range file{
			counts_temp := make(map[string]int)
			f,err:=os.Open(arg)
			if err!=nil{
				fmt.Fprintln(os.Stderr,"dup2:%v\n",err) //%v 以默认形式显示
			}
			countLines(f,counts_temp)
			for _,n:=range counts_temp{
				if n>1{
					fmt.Printf("%s\n",arg)
				}
			}
			f.Close()
		}
	}
	// for line,n:=range counts{
	// 	if n>1{
	// 		fmt.Println("%d\t%s\n",n,line)
	// 	}
	// }

}

func countLines(f *os.File,counts map[string]int){ //这里不copy一份counts，而是引用
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}