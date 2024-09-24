package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	for _,url := range os.Args[1:]{
		res,err:=http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"Fetch:%v\n",err)
			os.Exit(1)
		}
		// b,err := io.ReadAll(res.Body)
		// io.Copy(os.Stdout,io.ReadAll(res.Body))
		io.Copy(os.Stdout,res.Body)
		// 不需要装下整个数据流的缓冲区，直接将数据流拷贝到标准输出
		// res.Body.Close()
		// if err != nil{
		// 	fmt.Fprintf(os.Stderr,"fetch reading %s:%v",url,err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s",b)
	}
}