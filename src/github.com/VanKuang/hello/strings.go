package main

import "fmt"

type Person struct {
	ID int64
	Name string
}

type Deal struct {
	Product string
	Price float64
}

func (p Person) String() string {
	return fmt.Sprintf("[id: %v, name: %v]", p.ID, p.Name)
}

func main()  {
	kobe := Person{1, "Kobe"}
	deal := Deal{"SPT", 10.9}

	fmt.Println(kobe, deal)
}
