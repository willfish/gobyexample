package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums)
	fmt.Println()
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	l := len(nums)

	fmt.Println("Length of nums is", l)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
