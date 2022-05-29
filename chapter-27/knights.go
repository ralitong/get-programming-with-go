package main

import "fmt"

type character struct {
	name string
	item
}

func (c character) String() string {
	return c.name
}

type item string

func (c *character) pickup(i item)  {
	fmt.Printf("%v picks up %v\n", c, i)
	c.item = i
}

func (c *character) give(to *character) {
	fmt.Printf("%v gives %v to %v\n",c, c.item, to)
	to.item = c.item
	c.item = item("nil")
}

func (c character) status() {
	fmt.Printf("%v has %v\n", c, c.item)
}

func main() {
	sword := item("sword")
	nilItem := item("nil")
	arthur := character{"Arthur", nilItem}
	knight := character{"Knight", nilItem}
	arthur.status()
	knight.status()
	arthur.pickup(sword)
	arthur.give(&knight)
	arthur.status()
	knight.status()
}