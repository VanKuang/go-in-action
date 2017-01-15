package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	s1 := s[1 : 5]
	fmt.Println(s1)

	s1[0] = 0
	fmt.Println(s1)
	fmt.Println(s)

	fmt.Println(s[:9])
	fmt.Println(s[1:])

	var s2 []int
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s2 == nil)

	b := make([]int, 5)
	fmt.Println(b, len(b), cap(b))

	c := make([]int, 0, 5)
	fmt.Println(c, len(c), cap(c))

	fmt.Println(c[:2])
	fmt.Println(c[2:5])

	fmt.Println(append(c, 1))

	for i, v := range s {
		fmt.Printf("s[%d] == %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow {
		fmt.Println(v)
	}
}
