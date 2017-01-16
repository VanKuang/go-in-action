package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if (f < 0) {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M()  {
	fmt.Println(t.S)
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var a Abser
	f := MyFloat(-1)
	v := Vertex{3, 4}

	// f implement Abser
	a = f
	fmt.Println(a.Abs())

	// &v implement Abser
	a = &v

	fmt.Println(a.Abs())

	var i I = T{"hello world"}
	i.M()

	// nil interface : panic
	//var i1 I
	//i1.M()

	var i2 interface{}
	fmt.Printf("%v, %T\n", i2, i2)

	i2 = 10
	fmt.Printf("%v, %T\n", i2, i2)

	i2 = "hello"
	fmt.Printf("%v, %T\n", i2, i2)

	var i3 interface{} = "hello"

	s := i3.(string)
	fmt.Println(s)

	s, ok := i3.(string)
	fmt.Printf("%v is string? %v\n", s, ok)

	s1, ok := i3.(int)
	fmt.Printf("%v is int? %v\n", s1, ok)

	// panic
	//f = i3.(float64)
	//fmt.Printf("%v is float? %v\n", f, ok)

	typeSwitch(1)
	typeSwitch(1.1)
	typeSwitch("hello")
	typeSwitch(T{"hello"})
}