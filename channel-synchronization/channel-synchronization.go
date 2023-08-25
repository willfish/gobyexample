package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// This example shows how to wait for multiple goroutines to finish.
// We create blocking receive for each of the goroutines.
// We do not need to do anything with the data we receive from the goroutines,
// but we must receive something before the program can exit.
func main() {
	done := make(chan bool, 3)
	for i := 0; i < 3; i++ {
		go worker(done)
	}

	for i := 0; i < 3; i++ {
		<-done
	}
}
