package main

import "fmt"

type Base struct {
	num int
}

func (b Base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.num)
}

type container struct {
	Base
	str string
}

func main() {
	co := container{
		Base: Base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.Base.num)

	fmt.Println("describe:", co.describe())

	type Describer interface {
		describe() string
	}

	var d Describer = co
	fmt.Println("Describer:", d.describe())
}
