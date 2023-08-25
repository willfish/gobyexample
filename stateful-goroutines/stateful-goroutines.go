package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var readOps uint64
var writeOps uint64

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func statefulHandler(reads <-chan readOp, writes <-chan writeOp) {
	var state = make(map[int]int)
	for {
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.val
			write.resp <- true
		}
	}
}

func readOperationGenerator(reads chan<- readOp) {
	for {
		read := readOp{
			key:  rand.Intn(5),
			resp: make(chan int),
		}
		reads <- read
		<-read.resp
		atomic.AddUint64(&readOps, 1)
		time.Sleep(time.Millisecond)
	}
}

func writeOperationGenerator(writes chan<- writeOp) {
	for {
		write := writeOp{
			key:  rand.Intn(5),
			val:  rand.Intn(100),
			resp: make(chan bool),
		}
		writes <- write
		<-write.resp
		atomic.AddUint64(&writeOps, 1)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go statefulHandler(reads, writes)
	for r := 0; r < 100; r++ {
		go readOperationGenerator(reads)
	}
	for w := 0; w < 10; w++ {
		go writeOperationGenerator(writes)
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
