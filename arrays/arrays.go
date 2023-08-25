package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("dcl:", b)

	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(twoD[i][j])
		}
	}
}
