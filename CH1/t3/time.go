package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    // 测试字符串拼接的性能
    start := time.Now()
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = "*"
    }
    duration := time.Since(start)
    fmt.Printf("字符串拼接耗时: %v\n", duration)

    // 测试 strings.Join 的性能
    start = time.Now()
    joined := strings.Join(os.Args[1:], "*")
    duration = time.Since(start)
    fmt.Printf("strings.Join 耗时: %v\n", duration)

    // 打印结果以防止编译器优化
    fmt.Println(s)
    fmt.Println(joined)
}