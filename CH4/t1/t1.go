package main

import (
	"crypto/sha256"
	"fmt"
)

// 统计sha256散列中不同的位数

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	cnt := 0
	for i := 0; i < 32; i++ {
		xor := c1[i] ^ c2[i]
		for xor != 0 {
			cnt += int(xor & 1)
			xor >>= 1
		}
	}
	fmt.Println(cnt)
}
