package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 创建一个字符串通道

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 异步调用fetch函数，for执行完之后有多个fetch在一起执行，而不是一个fetch执行完再去执行另外一个fetch
	} // 异步的fetch函数和main函数通过ch字符串通道进行通信
	for range os.Args[1:] {
		fmt.Println(<-ch) //将ch的输出输入到println里
	}
	fmt.Printf("%.2fs elased\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	url = "http://" + url
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) // io.Discard 丢弃输出流
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
