package main

import (
	"fmt"
	"sync"
)

func streamNumbers(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	numberChannel := streamNumbers(2, 3)
	squaredChannel := squareNumber(squareNumber(numberChannel))
	fmt.Println(<-squaredChannel)
	fmt.Println(<-squaredChannel)

	numberChannel = streamNumbers(2, 3)
	// fan out the numbers in a pipeline (basically do lots of work in parallel - conceived as being heavier tasks than squaring a number)
	squaredChannel1 := squareNumber(numberChannel)
	squaredChannel2 := squareNumber(numberChannel)

	// fan in the squared numbers (the heavier work is done and we need to zip the results back together)
	for n := range merge(squaredChannel1, squaredChannel2) {
		fmt.Println(n)
	}
}
