package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var (
		indicator bool = true
		maxInt uint64 = 1 << 64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", indicator, indicator)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
