package main

import (
	"time"
	"fmt"
)

type MyError struct {
	Time time.Time
	Desc string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.Time, e.Desc)
}

func run() error {
	return &MyError{
		time.Now(),
		"something abnormal happened",
	}
}

func main()  {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
