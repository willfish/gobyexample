package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("init:", s, s == nil, len(s) == 0)
	fmt.Println("init:", s[0], s[1], s[2])
	fmt.Println("init:", len(s), cap(s))

	fmt.Println("init:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	fmt.Println("append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
