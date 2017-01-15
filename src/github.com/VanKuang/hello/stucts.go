package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 2
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	v1 := Vertex{X: 1}
	v2 := Vertex{}
	p = &Vertex{1, 2}
	fmt.Println(v1, v2, p)
}
