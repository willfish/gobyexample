package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	}

	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	} else if num < 10 {
		fmt.Println("Number has 1 digit")
	} else {
		fmt.Println("Number has multiple digits")
	}
}
