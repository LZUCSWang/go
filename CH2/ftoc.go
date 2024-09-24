package main

import "fmt"

func main() {
	const freezF, boilF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezF, fToC(freezF))
	fmt.Printf("%g F = %g C\n", boilF, fToC(boilF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
