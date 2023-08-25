package main

import "fmt"

// MapKeys returns a slice containing the keys of the map m.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(MapKeys(m))

	m2 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(MapKeys(m2))

	m3 := map[float64]bool{1.1: true, 2.2: false, 3.3: true}
	fmt.Println(MapKeys(m3))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(24)
	lst.Push(24)
	fmt.Println("list:", lst.GetAll())
}
