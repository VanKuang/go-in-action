package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = make(map[string]Vertex)
	m["China"] = Vertex{
		33.222, -99.88,
	}
	fmt.Println(m)
	fmt.Println(m["China"])
	fmt.Println(m["UK"])

	var m1 = map[string]Vertex {
		"USA": {111.22, 33.22},
		"HK": {23.11, 33.22},
	}
	fmt.Println(m1)

	m2 := make(map[int]string)
	m2[1] = "Kobe"
	fmt.Println(m2[1])
	m2[1] = "James"
	fmt.Println(m2[1])
	delete(m2, 1)
	fmt.Println(m2[1])
	v, ok := m2[1]
	fmt.Printf("Value:%v, exist:%v", v, ok)
}
