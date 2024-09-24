package main

import (
	"fmt"
	"math"
)

const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	fmt.Println(f == f+2)    // “false"
	fmt.Println(f == f+3)    // "false"

	fmt.Printf("%g %g %g\n", e, Avogadro, Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	fmt.Println(math.IsNaN(z), math.IsNaN(z/z), math.IsNaN(1/z))

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan != nan) // nan的比较总不成立，除了!=，它与==相反

}

func compute() (value float64, ok bool) {
	//
	var failed bool
	var result float64
	if failed {
		return 0, false
	}
	return result, true
}
