package main

import "fmt"

func main() {
	i := 42;
	f := 3.1415
	z := 2 + 3i
	fmt.Printf("Type: %T, Value: %v\n", i, i)
	fmt.Printf("Type: %T, Value: %v\n", f, f)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
