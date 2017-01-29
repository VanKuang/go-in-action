package main

import (
	"sync"
	"time"
	"fmt"
)

type Counter struct {
	count int64
}

type SafeCounter struct {
	count int64
	mux   sync.Mutex
}

func (counter *Counter) inc() {
	counter.count++
}

func (counter *SafeCounter) inc() {
	counter.mux.Lock()

	counter.count++

	counter.mux.Unlock()
}

func (counter *SafeCounter) value() int64 {
	counter.mux.Lock()

	defer counter.mux.Unlock()

	return counter.count
}

func main() {
	counter := Counter{0}
	safeCounter := SafeCounter{count: 0}

	for i := 0; i < 1000; i++ {
		go counter.inc()
		go safeCounter.inc()
	}

	time.Sleep(time.Second)

	fmt.Println(counter.count)
	fmt.Println(safeCounter.value())
}

