package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs and channel closed")
				fmt.Println(j)
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	fmt.Println("sent all jobs")

	fmt.Println("closing jobs queue")
	close(jobs)

  <-done
}
