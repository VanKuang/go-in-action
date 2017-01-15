package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	var cLang, java, goLang = false, true, true
	fmt.Println(i, j, k, cLang, java, goLang)

	apple, google, tcl := "good", "good", "bad"
	fmt.Println(apple, google, tcl)
}
