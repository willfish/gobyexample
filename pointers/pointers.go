package main

import "fmt"

func zeroval(ival int) {
	fmt.Println("zeroval:", &ival)
	ival = 0
}

func zeroptr(iptr *int) { // *int means that the function takes an integer pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", &i) // The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("initial:", i)

	zeroval(i) // zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
	fmt.Println("zeroval:", i)

	zeroptr(&i) // The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // Pointers can be printed too.
}
