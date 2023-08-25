package main

import "os"

func thisPanics() int {
	if true {
		panic("a function panics")
		return 1
	} else {
		return 0
	}
}

func main() {
	// thisPanics()
	// panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
