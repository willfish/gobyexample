package main

import "fmt"

func main() {
  // make a channel which will receive strings via a pipe
	messages := make(chan string)

  // asynchronously send a message to the pipe
	go func() {
		messages <- "ping" // blocking
	}()

	msg := <-messages // blocking
	newMsg := <-messages // blocking

	fmt.Println(msg)
	fmt.Println(newMsg)
}
