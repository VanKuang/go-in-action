package main

import "fmt"

func main() {
	defer fmt.Println("defer statement")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Function end")
}
