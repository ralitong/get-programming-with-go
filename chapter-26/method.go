package main

import "fmt"

type person struct {
	name, superpower string
	age int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	rebecca := person{
		"Rebecca",
		"imagination",
		14,
	}

	rebecca.birthday()
	fmt.Printf("%+v\n", rebecca)

	terry := person{
		"Terry",
		"Patience",
		15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := person{
		"Nathan",
		"Kindness",
		17,
	}

	nathan.birthday()
	fmt.Printf("%+v\n", nathan)
}