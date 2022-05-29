package main

import "fmt"

type person struct {
	name, superpower string
	age int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		"Rebecca",
		"imagination",
		14,
	}

	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
}