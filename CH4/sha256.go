package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	x := []byte("x")
	X := []byte("X")
	var x32 [32]byte
	var X32 [32]byte
	copy(x32[:], x)
	copy(X32[:], X)
	fmt.Println(x, X)
	zero(&x32)
	zero1(&X32)
	fmt.Println(x32, X32)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero1(ptr *[32]byte) {
	*ptr = [32]byte{}
}
