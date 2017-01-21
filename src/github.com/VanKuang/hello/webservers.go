package main

import (
	"net/http"
	"fmt"
	"log"
)

type Hello struct {
}

type E string

func (e E) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "error")
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello")
}

func main()  {
	var hello Hello

	var handler http.Handler
	e := E("ERROR")
	handler = e

	http.Handle("/error", handler)
	http.Handle("/", hello)
	err := http.ListenAndServe("localhost:4000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
