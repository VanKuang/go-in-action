package main

import "fmt"

func main() {
	fmt.Println(derivePositiveNumber(1))
	fmt.Println(derivePositiveNumber(-1))
}

func derivePositiveNumber(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}
