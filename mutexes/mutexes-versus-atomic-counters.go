package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int {
	return int(atomic.LoadInt64(&c.value))
}

func main() {
	var counterA, counterB Counter

	var wg sync.WaitGroup

	doIncrement := func(c *Counter, n int) {
		for i := 0; i < n; i++ {
			c.Inc()
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(&counterA, 10000)
	go doIncrement(&counterA, 10000)
	go doIncrement(&counterB, 10000)

	wg.Wait()

	counters := map[string]int{
		"a": counterA.Value(),
		"b": counterB.Value(),
	}

	fmt.Println(counters)
}
