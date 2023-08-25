package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func fib(n int) int {
	memo := make(map[int]int)
	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	// if value exists in memo, return it
	if val, exists := memo[n]; exists {
		return val
	}

	// base case
	if n < 2 {
		return n
	}

	// compute and store in memo before returning
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(fact(7))

	fmt.Println(fib(50))
}
