package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0;
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func putToChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i

		time.Sleep(1 * time.Second)
	}

	close(c)
}

func main() {
	a := []int{1, 3, 5, 7, 9}

	c := make(chan int)

	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x + y)

	c1 := make(chan int)

	go putToChannel(c1)

	for i := range c1 {
		fmt.Println(i)
	}
}