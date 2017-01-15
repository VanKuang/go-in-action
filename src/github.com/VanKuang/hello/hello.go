package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"github.com/VanKuang/utils"
)

func main() {
	fmt.Println("Hello, GO!")
	fmt.Println("Now: ", time.Now())
	fmt.Println("Square: ", math.Sqrt(9))
	fmt.Println("Number: ", rand.Intn(10))

	fmt.Println("1+2=", utils.Add(1, 2))
	fmt.Println(utils.Split(10))
	fmt.Println(utils.Reverse("!OG, olleH"))
	fmt.Println(utils.Swap("Hello", "world"))
}