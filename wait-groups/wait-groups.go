package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		workerId := i // avoid variable sharing (i is always the same and would be shared by each goroutine)

		go func() {
			defer wg.Done()
			worker(workerId)
		}()
	}

  // Blocks until all work is done
	wg.Wait()
}
