package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (f MyFloat) abs() float64  {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())

	v.scale(10)
	fmt.Println(v.abs())

	scale(&v, 10)
	fmt.Println(v.abs())
}
